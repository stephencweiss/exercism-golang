// Package strand is an interpreter of DNA
package strand

// ToRNA takes a DNA string and converts it to RNA
func ToRNA(dna string) string {
	rna := ""
	dnaToRna := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	for index := 0; index < len(dna); index++ {
		char := string(dna[index])
		if rnaChar, ok := dnaToRna[char]; ok {
			rna = rna + rnaChar
		} else {
			panic("Unknown DNA")
		}
	}
	return rna
}
