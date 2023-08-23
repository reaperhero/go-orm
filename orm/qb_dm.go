// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"fmt"
	"strconv"
	"strings"
)

// dmCommaSpace is the separation
const dmCommaSpace = ", "
const dmCommaFiled = "\""

// DmQueryBuilder is the SQL build
type DmQueryBuilder struct {
	tokens []string
}

// Select will join the fields
// Select("user.name","profile.age")
func (qb *DmQueryBuilder) Select(fields ...string) QueryBuilder {
	for i, field := range fields {
		if strings.Contains(field, "*") && len(strings.Split(field, ".")) == 2 {
			fields[i] = dmCommaFiled + strings.Split(field, ".")[0] + dmCommaFiled + "." + strings.Split(field, ".")[1]
			continue
		}
		if strings.Contains(field, "*") && len(strings.Split(field, ".")) == 1 {
			continue
		}
		if len(strings.Split(field, ".")) == 2 {
			fields[i] = dmCommaFiled + strings.Split(field, ".")[0] + dmCommaFiled + "." + dmCommaFiled + strings.Split(field, ".")[1] + dmCommaFiled
			continue
		} else {
			fields[i] = dmCommaFiled + field + dmCommaFiled
		}
	}
	qb.tokens = append(qb.tokens, "SELECT", strings.Join(fields, dmCommaSpace))
	return qb
}

// ForUpdate add the FOR UPDATE clause
func (qb *DmQueryBuilder) ForUpdate() QueryBuilder {
	qb.tokens = append(qb.tokens, "FOR UPDATE")
	return qb
}

// From join the tables
// From("deploy_cluster_host_rel")
func (qb *DmQueryBuilder) From(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "FROM", dmCommaFiled+strings.Join(tables, dmCommaSpace)+dmCommaFiled)
	return qb
}

// InnerJoin INNER JOIN the table
// InnerJoin("deploy_host")
func (qb *DmQueryBuilder) InnerJoin(table string) QueryBuilder {
	qb.tokens = append(qb.tokens, "INNER JOIN", dmCommaFiled+table+dmCommaFiled)
	return qb
}

// LeftJoin LEFT JOIN the table
// LeftJoin("deploy_host")
func (qb *DmQueryBuilder) LeftJoin(table string) QueryBuilder {
	qb.tokens = append(qb.tokens, "LEFT JOIN", dmCommaFiled+table+dmCommaFiled)
	return qb
}

// RightJoin RIGHT JOIN the table
// InnerJoin("deploy_host")
func (qb *DmQueryBuilder) RightJoin(table string) QueryBuilder {
	qb.tokens = append(qb.tokens, "RIGHT JOIN", dmCommaFiled+table+dmCommaFiled)
	return qb
}

// On join with on cond
// On(`"deploy_cluster_host_rel"."sid" = "deploy_host"."sid"`)
func (qb *DmQueryBuilder) On(cond string) QueryBuilder {
	splits := strings.Split(cond, "=")
	if len(splits) != 2 {
		return qb
	}
	qb.tokens = append(qb.tokens, "ON")
	for i, split := range splits {
		left := fieldCond(split)
		qb.tokens = append(qb.tokens, left)
		if i == 0 {
			qb.tokens = append(qb.tokens, "=")
		}
	}
	return qb
}

// Where join the Where cond
// Where("deploy_host.is_deleted = 1")
func (qb *DmQueryBuilder) Where(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "WHERE")

	operatr := condOperator(cond)
	if operatr == "" {
		return qb
	}
	splits := Split(cond, operatr)
	left := fieldCond(splits[0])
	right := fieldCondRight(splits[1])
	qb.tokens = append(qb.tokens, left, operatr, right)
	return qb
}

// And join the and cond
// And("deploy_host.is_deleted = 0")
func (qb *DmQueryBuilder) And(cond string) QueryBuilder {

	operatr := condOperator(cond)
	if operatr == "" {
		return qb
	}
	qb.tokens = append(qb.tokens, "AND")

	splits := Split(cond, operatr)
	left := fieldCond(splits[0])
	right := fieldCondRight(splits[1])
	qb.tokens = append(qb.tokens, left, operatr, right)
	return qb
}

// Or join the or cond
func (qb *DmQueryBuilder) Or(cond string) QueryBuilder {
	operatr := condOperator(cond)
	if operatr == "" {
		return qb
	}
	qb.tokens = append(qb.tokens, "OR")

	splits := Split(cond, operatr)
	left := fieldCond(splits[0])
	right := fieldCondRight(splits[1])
	qb.tokens = append(qb.tokens, left, operatr, right)
	return qb
}

// In join the IN (vals)
func (qb *DmQueryBuilder) In(vals ...string) QueryBuilder {
	var data []string
	for _, val := range vals {
		if strings.Trim(val, " ") == "?" {
			data = append(data, val)
			continue
		}
		_, err := strconv.Atoi(val)
		if err != nil {
			data = append(data, val, "'"+val+"'")
		} else {
			data = append(data, val, val)
		}
	}
	qb.tokens = append(qb.tokens, "IN", "(", strings.Join(data, dmCommaSpace), ")")
	return qb
}

// OrderBy join the Order by fields
// OrderBy("id")
func (qb *DmQueryBuilder) OrderBy(fields ...string) QueryBuilder {
	for i, field := range fields {
		fields[i] = fieldCond(field)
	}
	qb.tokens = append(qb.tokens, "ORDER BY", strings.Join(fields, dmCommaSpace))
	return qb
}

// Asc join the asc
func (qb *DmQueryBuilder) Asc() QueryBuilder {
	qb.tokens = append(qb.tokens, "ASC")
	return qb
}

// Desc join the desc
func (qb *DmQueryBuilder) Desc() QueryBuilder {
	qb.tokens = append(qb.tokens, "DESC")
	return qb
}

// Limit join the limit num
func (qb *DmQueryBuilder) Limit(limit int) QueryBuilder {
	qb.tokens = append(qb.tokens, "LIMIT", strconv.Itoa(limit))
	return qb
}

// Offset join the offset num
func (qb *DmQueryBuilder) Offset(offset int) QueryBuilder {
	qb.tokens = append(qb.tokens, "OFFSET", strconv.Itoa(offset))
	return qb
}

// GroupBy join the Group by fields
func (qb *DmQueryBuilder) GroupBy(fields ...string) QueryBuilder {
	for i, field := range fields {
		fields[i] = fieldCond(field)
	}
	qb.tokens = append(qb.tokens, "GROUP BY", strings.Join(fields, dmCommaSpace))
	return qb
}

// Having join the Having cond
func (qb *DmQueryBuilder) Having(cond string) QueryBuilder {
	qb.tokens = append(qb.tokens, "HAVING")

	operatr := condOperator(cond)
	if operatr == "" {
		return qb
	}
	splits := Split(cond, operatr)
	left := fieldCond(splits[0])
	right := fieldCondRight(splits[1])
	qb.tokens = append(qb.tokens, left, operatr, right)
	return qb
}

// Update join the update table
func (qb *DmQueryBuilder) Update(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "UPDATE", dmCommaFiled+strings.Join(tables, dmCommaFiled+dmCommaSpace+dmCommaFiled)+dmCommaFiled)
	return qb
}

// Set join the set kv
func (qb *DmQueryBuilder) Set(kv ...string) QueryBuilder {
	var data []string
	for _, s := range kv {
		operatr := condOperator(s)
		splits := strings.Split(s, operatr)
		left := fieldCond(splits[0])
		right := fieldCondRight(splits[1])
		data = append(data, left+operatr+right)
	}
	qb.tokens = append(qb.tokens, "SET", strings.Join(data, dmCommaSpace))
	return qb
}

// Delete join the Delete tables
func (qb *DmQueryBuilder) Delete(tables ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "DELETE")
	if len(tables) != 0 {
		qb.tokens = append(qb.tokens, dmCommaFiled+strings.Join(tables, dmCommaFiled+dmCommaSpace+dmCommaFiled)+dmCommaFiled)
	}
	return qb
}

// InsertInto join the insert SQL
func (qb *DmQueryBuilder) InsertInto(table string, fields ...string) QueryBuilder {
	qb.tokens = append(qb.tokens, "INSERT INTO", dmCommaFiled+table+dmCommaFiled)
	if len(fields) != 0 {
		fieldsStr :=dmCommaFiled + strings.Join(fields, dmCommaFiled+dmCommaSpace+dmCommaFiled)+dmCommaFiled
		qb.tokens = append(qb.tokens, "(", fieldsStr, ")")
	}
	return qb
}

// Values join the Values(vals)
func (qb *DmQueryBuilder) Values(vals ...string) QueryBuilder {
	var data []string
	for _, val := range vals {
		right := fieldCondRight(val)
		data = append(data,right)
	}
	valsStr := strings.Join(data, dmCommaSpace)
	qb.tokens = append(qb.tokens, "VALUES", "(", valsStr, ")")
	return qb
}

// Subquery join the sub as alias
func (qb *DmQueryBuilder) Subquery(sub string, alias string) string {
	return fmt.Sprintf("(%s) AS %s", sub, alias)
}

// String join all tokens
func (qb *DmQueryBuilder) String() string {
	s := strings.Join(qb.tokens, " ")
	qb.tokens = qb.tokens[:0]
	return s
}

func Split(source string, split string) []string {
	data := strings.Split(source, split)
	if len(data) == 0 && data[0] == "" {
		return nil
	}
	return data
}

func fieldCond(cond string) string {
	var data string
	splits := strings.Split(cond, ".")
	if len(splits) == 2 {
		data = dmCommaFiled +strings.Trim( splits[0]," ") + dmCommaFiled + "." + dmCommaFiled + strings.Trim(splits[1]," ") + dmCommaFiled
	} else {
		data = dmCommaFiled + strings.Trim(splits[0]," ") + dmCommaFiled
	}
	return data
}

// 处理条件右边的语句
func fieldCondRight(value string) string {
	var data string
	if strings.Trim(value, " ") == "?" {
		return value
	}
	if strings.HasPrefix(value, "'") {
		return value
	}
	_, err := strconv.Atoi(value)
	if err != nil {
		return value
	}
	splits := strings.Split(value, ".")
	if len(splits) == 2 {
		data = dmCommaFiled + strings.Trim(splits[0]," ") + dmCommaFiled + "." + dmCommaFiled + strings.Trim(splits[1]," ") + dmCommaFiled
	} else {
		data = dmCommaFiled + strings.Trim(splits[0]," ") + dmCommaFiled
	}
	return data
}

var operatorKeys = []byte{'>', '=', '<', '!'}

func bytesContain(source []byte, search byte) bool {
	for _, b := range source {
		if b == search {
			return true
		}
	}
	return false
}

func condOperator(cond string) string {
	var first int = -1
	for i, i2 := range cond {
		if first != -1 && !bytesContain(operatorKeys, byte(i2)) {
			return string(cond[first:i])
		}
		if first == -1 && bytesContain(operatorKeys, byte(i2)) {
			first = i
		}
	}
	return ""
}
