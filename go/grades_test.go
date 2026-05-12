package main

import "testing"

func ptr(v int) *int { return &v }

func sampleTable() []GradeRow {
	return []GradeRow{
		{Grade: 1, Amount: 58000, Lower: nil, Upper: ptr(63000)},
		{Grade: 2, Amount: 68000, Lower: ptr(63000), Upper: ptr(73000)},
		{Grade: 3, Amount: 78000, Lower: ptr(73000), Upper: ptr(83000)},
		{Grade: 21, KenpoGrade: ptr(18), Amount: 280000, Lower: ptr(270000), Upper: ptr(290000)},
		{Grade: 22, KenpoGrade: ptr(19), Amount: 300000, Lower: ptr(290000), Upper: ptr(310000)},
		{Grade: 23, KenpoGrade: ptr(20), Amount: 320000, Lower: ptr(310000), Upper: ptr(330000)},
		{Grade: 50, Amount: 1390000, Lower: ptr(1355000), Upper: nil},
	}
}

func TestResolve(t *testing.T) {
	tbl := sampleTable()
	cases := []struct {
		avg  int
		want int
	}{
		{50000, 58000},
		{62999, 58000},
		{63000, 68000},
		{63001, 68000},
		{72999, 68000},
		{73000, 78000},
		{289999, 280000},
		{290000, 300000},
		{300000, 300000},
		{309999, 300000},
		{310000, 320000},
		{2000000, 1390000},
	}
	for _, c := range cases {
		row, err := Resolve(tbl, c.avg)
		if err != nil {
			t.Fatalf("avg=%d: %v", c.avg, err)
		}
		if row.Amount != c.want {
			t.Fatalf("avg=%d -> %d, want %d", c.avg, row.Amount, c.want)
		}
	}
}

func TestNextThreshold(t *testing.T) {
	tbl := sampleTable()
	row, _ := Resolve(tbl, 300000)
	nt := NextThreshold(tbl, row)
	if nt == nil || *nt != 310000 {
		t.Fatalf("next threshold = %v, want 310000", nt)
	}

	top, _ := Resolve(tbl, 2000000)
	if NextThreshold(tbl, top) != nil {
		t.Fatalf("top grade should have nil next threshold")
	}
}

func TestJudge(t *testing.T) {
	tbl := sampleTable()
	j, err := Judge(tbl, 300000)
	if err != nil {
		t.Fatal(err)
	}
	if j.StandardAmount != 300000 || j.Pension != 22 || *j.KenpoGrade != 19 {
		t.Fatalf("unexpected judgement: %+v", j)
	}
	if j.NextThreshold == nil || *j.NextThreshold != 310000 {
		t.Fatalf("next threshold = %v", j.NextThreshold)
	}
}
