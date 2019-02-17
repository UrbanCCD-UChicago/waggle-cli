package descriptions

// ListNodes :
const ListNodes = `Lists node information stored in the local cache

EXAMPLES
========

$ waggle list-nodes
0000001e06200335	THS1 	10001	Testing Rig 1                 	Surya
0000001E06107CC5	W00D 	10018	NU Uglybox                    	Tuley Park Court Yard
0000001E06107E4C	W002 	10019	None
0000001e06107d7f	W007 	10050	NULL                          	NULL
0000001E061089FA	011  	10027	AoT Chicago (T) - Returned    	TCS 4302,01/18/2018

$ waggle list-nodes --alive-only
0000001E06107CC5	W00D 	10018	NU Uglybox                    	Tuley Park Court Yard
0000001E06107E4C	W002 	10019	None
0000001e06107d7f	W007 	10050	NULL                          	NULL
0000001E06107D97	Ret_waggle_006	10032	NU-CBG Uglybox #1             	CBG
0000001E06200367	W008 	10042	Tokyo Univ. of Tech           	Tokyo, Japan

$ waggle list-nodes --format csv
"0000001e06200335","THS1","10001","Testing Rig 1","Surya"
"0000001E06107CC5","W00D","10018","NU Uglybox","Tuley Park Court Yard"
"0000001E06107E4C","W002","10019","None",""
"0000001e06107d7f","W007","10050","NULL","NULL"
"0000001E061089FA","011","10027","AoT Chicago (T) - Returned","TCS 4302,01/18/2018"

$ waggle list-nodes --alive-only --format csv
"0000001E06107CC5","W00D","10018","NU Uglybox","Tuley Park Court Yard"
"0000001E06107E4C","W002","10019","None",""
"0000001e06107d7f","W007","10050","NULL","NULL"
"0000001E06107D97","Ret_waggle_006","10032","NU-CBG Uglybox #1","CBG"
"0000001E06200367","W008","10042","Tokyo Univ. of Tech","Tokyo, Japan"

CAVEATS
=======

* This will only work on an admin computer.
`
