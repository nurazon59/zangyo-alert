module Domain.Grades

// 等級表の純粋ロジック。CSV読み込みは外（Io層）でやり、ここは値として受け取る。
//
// 型:
//   GradeRow   : CSV1行分 (Grade, KenpoGrade option, Amount, Lower option, Upper option)
//   GradeTable : GradeRow list
//
// 提供する関数:
//   resolve : GradeTable -> averageAmount: decimal -> GradeRow
//     Lower < averageAmount < Upper の行を返す
//     既存 Grades.fs の resolveGrades のロジックをここに移す
//
//   nextThreshold : GradeTable -> currentGrade: int -> decimal option
//     次等級の下限額。最上位なら None
//
//   judge : GradeTable -> averageAmount: decimal -> GradeJudgement
//     resolve + nextThreshold を組み合わせて Domain.Types.GradeJudgement を組み立てる
