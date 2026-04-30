open System.IO
open System.Text

let text = "testdata/sample.txt" |> File.ReadAllText |> printfn "%s"
