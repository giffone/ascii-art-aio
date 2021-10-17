// package structure strores all data structures
package structure

type Art struct {
	Args           []string // all arguments
	LenArgs        int      // lenght of arguments
	WarningMessage string   // text for error
	TrashArgs      []string // for deleted arguments
	Alphabet       Alphabet
	Flag           Flag
	T              T
	Text           Text
	Font           Font
	Output         Output
	Read           Read
}

type Alphabet struct {
	Sourse   []string
	Rune     map[rune][]string // sorted in ascii code
	RuneLen  map[rune]int      // save lenght of each symbol
	Letter   rune              // current letter
	Coloring bool              // if need to color
	Space    int               // width of rune " " (32)
	Reverse
}

type Text struct {
	Rune    []rune // user text
	Len     []int  // lenght of each symbol
	Count   int    // count letters
	Width   []int  // width words in ascii-ard in each line
	Spaces  []int  // count spaces in ech line
	Symbols bool   // checks if there are symbols exept space
}

type Output struct {
	Final  []string // make output
	Spaces int      // count spaces between words for justify
	Index  int      // lines in output 8+8+8+...
}

type Font struct {
	Value   string
	Founded bool
}

// T - size of terminal window
type T struct {
	Size   []int // height, width
	Height int
	Width  int
}

type Flag struct {
	Output  string
	Reverse string
	Align   string
	Color   Color
}

type Color struct {
	Value          string // red, green, blue and etc
	MethodBy       string // by indexes or letters
	MethodColoring string // method coloring: and (a-d), or (a,b,c), none, all
	Case           []byte
	Case1          string // color start
	Case2          string // color end (reset)
	BySymbol       BySymbol
	ByIndex        ByIndex
}

type BySymbol struct {
	Range  []string // to save letters (a-d), (a,b,c)
	Range1 rune     // first letter
	Range2 rune     // last letter
}

type ByIndex struct {
	Range    []int // to save indexes [1:4], [0,3,8]
	Range1   int   // first index
	Range2   int   // last index
	MaxIndex bool  // lenght of all letters in arguments (to make max index)
}

// Scan local structure for bufio.Scanner
type Scan struct {
	Slice      []byte // scanned line (scan line by line) in rune
	LLine      int    // lenght of scanned line
	SymbolR    rune   // current symbol in rune
	SymbolB    byte   // current symbol in byte
	Height     int    // height of each letter
	Code       int    // number of possition in ascii
	Lines      int    // lines in txt file
	Start      bool
	IndexStart int
	IndexEnd   int
	CountSpace int
}

// Resize local structure for function SplitArgs
type Resize struct {
	Add     bool // add or don't newline
	LSymbol int  // lenght of current symbol
	Width   int  // calculate width between newlines
	Spaces  int  // count spases
}

type Reverse struct {
	BigData map[int]map[int]byte // finished data of keys
	Fonts   []string
}

type Read struct {
	File      []byte //
	Width     []int
	Tmp       []byte
	TmpMatrix []byte //
	Answer    []byte
}

// --output=<fileName.txt>
// --color=<color>
// --reverse=<fileName>
// --align=<type>
