Title: Rewarding Contributors for {{.Interval.Start.Format "January 2006"}}

Content:

1. Object the confirmation of the contributed time by Contributors to teledisko DAO from {{.Interval.Start.Format "02.01.2006"}} to {{.Interval.EndInclude.Format "02.01.2006"}} or the minting of the corresponding number of tokens to them in the following manner:{{range $a := .TokenAllocations}}
    1. {{$a.Name}}, {{printf "%.2f" $a.HoursAmount}} hours, {{printf "%.2f" $a.TokenAmount}} tokens;{{end}}