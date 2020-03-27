package fn

import (
	"../config"
	"fmt"
	// "errors"
	"strings"
)

type Sql struct {
	Fields   []string
	Table    string
	Action   config.SqlAction
	WhereStr string
	ToInsertFields []string
	ToInsertVals [][]string
}

func (x *Sql) Select(fields []string) *Sql {
	x.Fields = fields
	//todo make const string
	x.Action = config.SelectType
	return x
}

func (x *Sql) Insert(fields []string, vals [][]string) *Sql {
	x.ToInsertFields = fields
	x.ToInsertVals = vals

	//todo make const string
	x.Action = config.InsertType
	return x
}

func (x *Sql) From(tbName string) *Sql {
	x.Table = tbName
	return x
}

func (x *Sql) Where(s string) *Sql {
	x.WhereStr = fmt.Sprintf(` WHERE %s`, s)
	return x
}

func (x *Sql) End() string {
	var d string
	var base string
	var segment1 string
	var segment2 string

	switch x.Action {
	case config.SelectType:
		segment1 = strings.Join(x.Fields, ",")
		segment2 = x.Table + ` ` + x.WhereStr
		d = fmt.Sprintf(`SELECT %v FROM %v`, segment1, segment2)
		// d = x.Action + ` ` + strings.Join(x.Fields, ",") + ` FROM ` + x.Table + ` ` + x.WhereStr
	case config.InsertType:
		base = `INSERT INTO %v VALUES %v`
		
		//["first", "second"]
		var fields string = strings.Join(x.ToInsertFields, ",")
		segment1 = x.Table + ` ` + `(` +fields+ `)`

		// [['hi', '3'], ['another', '5'], ['another', '5']] turn into ($1, $2),($3, $4),($5, $6)

		var sqlHolder int = 1
		
		var bigHolder []string
		var pairCnt int = len(x.ToInsertVals)
		var groupElemCnt int = len(x.ToInsertFields)
		for j := 0; j < pairCnt; j++ {

			var holder []string
			//pair == ['hi', '3']
			for i := 0; i < groupElemCnt; i++ {

				holder = append(holder, fmt.Sprintf(`$%v`, sqlHolder))
				sqlHolder++
				//holder == ["$1", "$2"]
			}
			
			var bracket string = `(`+ strings.Join(holder, ",")+`)`
			fmt.Println(`\nbracket: `, bracket)

			bigHolder = append(bigHolder, bracket)

			fmt.Println(`\nbigHolder: `, bigHolder)
		}

		segment2 = strings.Join(bigHolder, ",")

		fmt.Println(`\nfn.go INSERT case segment2: `, segment2)
		//($1,$2),($1,$2,$3,$4) --> Wrong

		d = fmt.Sprintf(base, segment1, segment2)
	default:
		fmt.Println("BAD x.Action")
	}

	//clean obj
	x.Fields = nil
	x.Table = config.EmptyStr
	x.Action = config.NoType
	x.WhereStr = config.EmptyStr
	x.ToInsertFields = nil
	x.ToInsertVals = nil

	return d
}

func Contains(arr []int, x int) bool {
	for _, a := range arr {
		if a == x {
			return true
		}
	}
	return false
}

func HandleErr(err error, place string) {
	var errNow string
	errNow = fmt.Sprintf(`ERRRR at (%v) err: (%v)`, place, err.Error())
	fmt.Println(errNow)
}

// func FlattenArr(bigarr [][]string) []string {
// 	var dout []string
// 	for _, arr := range bigarr {
// 		for _, elem := range arr {
// 			dout = append(dout, elem)
// 		}
// 	}
// 	return dout
// }