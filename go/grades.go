package main

import "fmt"

// Resolve は標準報酬月額表から該当等級行を返す。
// 区間の判定は「下限以上・上限未満」(Lower <= avg < Upper)。
// 協会けんぽの月額表は隣接区間が同値の境界（63000等）を共有しており、strict比較だと境界が落ちるため。
func Resolve(table []GradeRow, avg int) (GradeRow, error) {
	for _, row := range table {
		if row.Lower != nil && avg < *row.Lower {
			continue
		}
		if row.Upper != nil && avg >= *row.Upper {
			continue
		}
		return row, nil
	}
	return GradeRow{}, fmt.Errorf("no grade row matches average=%d", avg)
}

func NextThreshold(table []GradeRow, current GradeRow) *int {
	for _, row := range table {
		if row.Grade == current.Grade+1 {
			return row.Lower
		}
	}
	return nil
}

func Judge(table []GradeRow, avg int) (GradeJudgement, error) {
	row, err := Resolve(table, avg)
	if err != nil {
		return GradeJudgement{}, err
	}
	return GradeJudgement{
		StandardAmount: row.Amount,
		KenpoGrade:     row.KenpoGrade,
		Pension:        row.Grade,
		NextThreshold:  NextThreshold(table, row),
	}, nil
}
