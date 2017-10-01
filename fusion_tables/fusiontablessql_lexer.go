// Generated from FusionTablesSql.g4 by ANTLR 4.7.

package fusion_tables

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 80, 729,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9, 105, 4, 106,
	9, 106, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61,
	3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3,
	62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64,
	3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3,
	66, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71,
	3, 71, 3, 72, 3, 72, 3, 73, 6, 73, 608, 10, 73, 13, 73, 14, 73, 609, 3,
	73, 3, 73, 7, 73, 614, 10, 73, 12, 73, 14, 73, 617, 11, 73, 5, 73, 619,
	10, 73, 3, 73, 3, 73, 5, 73, 623, 10, 73, 3, 74, 3, 74, 5, 74, 627, 10,
	74, 3, 75, 6, 75, 630, 10, 75, 13, 75, 14, 75, 631, 3, 76, 3, 76, 3, 76,
	3, 76, 7, 76, 638, 10, 76, 12, 76, 14, 76, 641, 11, 76, 3, 76, 3, 76, 3,
	77, 3, 77, 3, 77, 3, 77, 7, 77, 649, 10, 77, 12, 77, 14, 77, 652, 11, 77,
	3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 78, 7, 78, 660, 10, 78, 12, 78, 14,
	78, 663, 11, 78, 3, 78, 3, 78, 3, 78, 5, 78, 668, 10, 78, 3, 78, 3, 78,
	3, 79, 3, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3,
	83, 3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 87, 3, 87, 3, 88,
	3, 88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3, 92, 3, 93, 3,
	93, 3, 94, 3, 94, 3, 95, 3, 95, 3, 96, 3, 96, 3, 97, 3, 97, 3, 98, 3, 98,
	3, 99, 3, 99, 3, 100, 3, 100, 3, 101, 3, 101, 3, 102, 3, 102, 3, 103, 3,
	103, 3, 104, 3, 104, 3, 105, 3, 105, 3, 106, 3, 106, 3, 661, 2, 107, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40,
	79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49,
	97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129,
	66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71, 141, 72, 143, 73, 145,
	74, 147, 75, 149, 76, 151, 77, 153, 78, 155, 79, 157, 80, 159, 2, 161,
	2, 163, 2, 165, 2, 167, 2, 169, 2, 171, 2, 173, 2, 175, 2, 177, 2, 179,
	2, 181, 2, 183, 2, 185, 2, 187, 2, 189, 2, 191, 2, 193, 2, 195, 2, 197,
	2, 199, 2, 201, 2, 203, 2, 205, 2, 207, 2, 209, 2, 211, 2, 3, 2, 34, 4,
	2, 45, 45, 47, 47, 7, 2, 47, 47, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2,
	41, 41, 4, 2, 12, 12, 15, 15, 5, 2, 11, 13, 15, 15, 34, 34, 3, 2, 50, 59,
	4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4,
	2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4,
	2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4,
	2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4,
	2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4,
	2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4,
	2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4,
	2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4,
	2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 2, 712, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89,
	3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2,
	97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2,
	2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111,
	3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2,
	2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3,
	2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2,
	133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2,
	2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147,
	3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2,
	2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 3, 213, 3, 2, 2, 2, 5, 215, 3,
	2, 2, 2, 7, 217, 3, 2, 2, 2, 9, 219, 3, 2, 2, 2, 11, 221, 3, 2, 2, 2, 13,
	227, 3, 2, 2, 2, 15, 231, 3, 2, 2, 2, 17, 234, 3, 2, 2, 2, 19, 237, 3,
	2, 2, 2, 21, 241, 3, 2, 2, 2, 23, 249, 3, 2, 2, 2, 25, 252, 3, 2, 2, 2,
	27, 260, 3, 2, 2, 2, 29, 265, 3, 2, 2, 2, 31, 272, 3, 2, 2, 2, 33, 279,
	3, 2, 2, 2, 35, 288, 3, 2, 2, 2, 37, 294, 3, 2, 2, 2, 39, 301, 3, 2, 2,
	2, 41, 308, 3, 2, 2, 2, 43, 313, 3, 2, 2, 2, 45, 322, 3, 2, 2, 2, 47, 327,
	3, 2, 2, 2, 49, 335, 3, 2, 2, 2, 51, 340, 3, 2, 2, 2, 53, 345, 3, 2, 2,
	2, 55, 350, 3, 2, 2, 2, 57, 356, 3, 2, 2, 2, 59, 363, 3, 2, 2, 2, 61, 372,
	3, 2, 2, 2, 63, 375, 3, 2, 2, 2, 65, 382, 3, 2, 2, 2, 67, 387, 3, 2, 2,
	2, 69, 392, 3, 2, 2, 2, 71, 399, 3, 2, 2, 2, 73, 404, 3, 2, 2, 2, 75, 409,
	3, 2, 2, 2, 77, 415, 3, 2, 2, 2, 79, 423, 3, 2, 2, 2, 81, 431, 3, 2, 2,
	2, 83, 439, 3, 2, 2, 2, 85, 443, 3, 2, 2, 2, 87, 449, 3, 2, 2, 2, 89, 452,
	3, 2, 2, 2, 91, 459, 3, 2, 2, 2, 93, 462, 3, 2, 2, 2, 95, 468, 3, 2, 2,
	2, 97, 474, 3, 2, 2, 2, 99, 484, 3, 2, 2, 2, 101, 491, 3, 2, 2, 2, 103,
	503, 3, 2, 2, 2, 105, 510, 3, 2, 2, 2, 107, 524, 3, 2, 2, 2, 109, 528,
	3, 2, 2, 2, 111, 532, 3, 2, 2, 2, 113, 537, 3, 2, 2, 2, 115, 544, 3, 2,
	2, 2, 117, 550, 3, 2, 2, 2, 119, 557, 3, 2, 2, 2, 121, 560, 3, 2, 2, 2,
	123, 567, 3, 2, 2, 2, 125, 574, 3, 2, 2, 2, 127, 579, 3, 2, 2, 2, 129,
	585, 3, 2, 2, 2, 131, 590, 3, 2, 2, 2, 133, 593, 3, 2, 2, 2, 135, 596,
	3, 2, 2, 2, 137, 598, 3, 2, 2, 2, 139, 600, 3, 2, 2, 2, 141, 602, 3, 2,
	2, 2, 143, 604, 3, 2, 2, 2, 145, 622, 3, 2, 2, 2, 147, 626, 3, 2, 2, 2,
	149, 629, 3, 2, 2, 2, 151, 633, 3, 2, 2, 2, 153, 644, 3, 2, 2, 2, 155,
	655, 3, 2, 2, 2, 157, 671, 3, 2, 2, 2, 159, 675, 3, 2, 2, 2, 161, 677,
	3, 2, 2, 2, 163, 679, 3, 2, 2, 2, 165, 681, 3, 2, 2, 2, 167, 683, 3, 2,
	2, 2, 169, 685, 3, 2, 2, 2, 171, 687, 3, 2, 2, 2, 173, 689, 3, 2, 2, 2,
	175, 691, 3, 2, 2, 2, 177, 693, 3, 2, 2, 2, 179, 695, 3, 2, 2, 2, 181,
	697, 3, 2, 2, 2, 183, 699, 3, 2, 2, 2, 185, 701, 3, 2, 2, 2, 187, 703,
	3, 2, 2, 2, 189, 705, 3, 2, 2, 2, 191, 707, 3, 2, 2, 2, 193, 709, 3, 2,
	2, 2, 195, 711, 3, 2, 2, 2, 197, 713, 3, 2, 2, 2, 199, 715, 3, 2, 2, 2,
	201, 717, 3, 2, 2, 2, 203, 719, 3, 2, 2, 2, 205, 721, 3, 2, 2, 2, 207,
	723, 3, 2, 2, 2, 209, 725, 3, 2, 2, 2, 211, 727, 3, 2, 2, 2, 213, 214,
	7, 61, 2, 2, 214, 4, 3, 2, 2, 2, 215, 216, 7, 44, 2, 2, 216, 6, 3, 2, 2,
	2, 217, 218, 7, 46, 2, 2, 218, 8, 3, 2, 2, 2, 219, 220, 7, 48, 2, 2, 220,
	10, 3, 2, 2, 2, 221, 222, 5, 161, 81, 2, 222, 223, 5, 183, 92, 2, 223,
	224, 5, 199, 100, 2, 224, 225, 5, 169, 85, 2, 225, 226, 5, 195, 98, 2,
	226, 12, 3, 2, 2, 2, 227, 228, 5, 161, 81, 2, 228, 229, 5, 187, 94, 2,
	229, 230, 5, 167, 84, 2, 230, 14, 3, 2, 2, 2, 231, 232, 5, 189, 95, 2,
	232, 233, 5, 195, 98, 2, 233, 16, 3, 2, 2, 2, 234, 235, 5, 161, 81, 2,
	235, 236, 5, 197, 99, 2, 236, 18, 3, 2, 2, 2, 237, 238, 5, 161, 81, 2,
	238, 239, 5, 197, 99, 2, 239, 240, 5, 165, 83, 2, 240, 20, 3, 2, 2, 2,
	241, 242, 5, 161, 81, 2, 242, 243, 5, 203, 102, 2, 243, 244, 5, 169, 85,
	2, 244, 245, 5, 195, 98, 2, 245, 246, 5, 161, 81, 2, 246, 247, 5, 173,
	87, 2, 247, 248, 5, 169, 85, 2, 248, 22, 3, 2, 2, 2, 249, 250, 5, 163,
	82, 2, 250, 251, 5, 209, 105, 2, 251, 24, 3, 2, 2, 2, 252, 253, 5, 163,
	82, 2, 253, 254, 5, 169, 85, 2, 254, 255, 5, 199, 100, 2, 255, 256, 5,
	205, 103, 2, 256, 257, 5, 169, 85, 2, 257, 258, 5, 169, 85, 2, 258, 259,
	5, 187, 94, 2, 259, 26, 3, 2, 2, 2, 260, 261, 5, 165, 83, 2, 261, 262,
	5, 161, 81, 2, 262, 263, 5, 197, 99, 2, 263, 264, 5, 169, 85, 2, 264, 28,
	3, 2, 2, 2, 265, 266, 5, 165, 83, 2, 266, 267, 5, 177, 89, 2, 267, 268,
	5, 195, 98, 2, 268, 269, 5, 165, 83, 2, 269, 270, 5, 183, 92, 2, 270, 271,
	5, 169, 85, 2, 271, 30, 3, 2, 2, 2, 272, 273, 5, 165, 83, 2, 273, 274,
	5, 189, 95, 2, 274, 275, 5, 183, 92, 2, 275, 276, 5, 201, 101, 2, 276,
	277, 5, 185, 93, 2, 277, 278, 5, 187, 94, 2, 278, 32, 3, 2, 2, 2, 279,
	280, 5, 165, 83, 2, 280, 281, 5, 189, 95, 2, 281, 282, 5, 187, 94, 2, 282,
	283, 5, 199, 100, 2, 283, 284, 5, 161, 81, 2, 284, 285, 5, 177, 89, 2,
	285, 286, 5, 187, 94, 2, 286, 287, 5, 197, 99, 2, 287, 34, 3, 2, 2, 2,
	288, 289, 5, 165, 83, 2, 289, 290, 5, 189, 95, 2, 290, 291, 5, 201, 101,
	2, 291, 292, 5, 187, 94, 2, 292, 293, 5, 199, 100, 2, 293, 36, 3, 2, 2,
	2, 294, 295, 5, 165, 83, 2, 295, 296, 5, 195, 98, 2, 296, 297, 5, 169,
	85, 2, 297, 298, 5, 161, 81, 2, 298, 299, 5, 199, 100, 2, 299, 300, 5,
	169, 85, 2, 300, 38, 3, 2, 2, 2, 301, 302, 5, 167, 84, 2, 302, 303, 5,
	169, 85, 2, 303, 304, 5, 183, 92, 2, 304, 305, 5, 169, 85, 2, 305, 306,
	5, 199, 100, 2, 306, 307, 5, 169, 85, 2, 307, 40, 3, 2, 2, 2, 308, 309,
	5, 167, 84, 2, 309, 310, 5, 169, 85, 2, 310, 311, 5, 197, 99, 2, 311, 312,
	5, 165, 83, 2, 312, 42, 3, 2, 2, 2, 313, 314, 5, 167, 84, 2, 314, 315,
	5, 169, 85, 2, 315, 316, 5, 197, 99, 2, 316, 317, 5, 165, 83, 2, 317, 318,
	5, 195, 98, 2, 318, 319, 5, 177, 89, 2, 319, 320, 5, 163, 82, 2, 320, 321,
	5, 169, 85, 2, 321, 44, 3, 2, 2, 2, 322, 323, 5, 167, 84, 2, 323, 324,
	5, 189, 95, 2, 324, 325, 5, 169, 85, 2, 325, 326, 5, 197, 99, 2, 326, 46,
	3, 2, 2, 2, 327, 328, 5, 165, 83, 2, 328, 329, 5, 189, 95, 2, 329, 330,
	5, 187, 94, 2, 330, 331, 5, 199, 100, 2, 331, 332, 5, 161, 81, 2, 332,
	333, 5, 177, 89, 2, 333, 334, 5, 187, 94, 2, 334, 48, 3, 2, 2, 2, 335,
	336, 5, 167, 84, 2, 336, 337, 5, 195, 98, 2, 337, 338, 5, 189, 95, 2, 338,
	339, 5, 191, 96, 2, 339, 50, 3, 2, 2, 2, 340, 341, 5, 169, 85, 2, 341,
	342, 5, 187, 94, 2, 342, 343, 5, 167, 84, 2, 343, 344, 5, 197, 99, 2, 344,
	52, 3, 2, 2, 2, 345, 346, 5, 171, 86, 2, 346, 347, 5, 195, 98, 2, 347,
	348, 5, 189, 95, 2, 348, 349, 5, 185, 93, 2, 349, 54, 3, 2, 2, 2, 350,
	351, 5, 173, 87, 2, 351, 352, 5, 195, 98, 2, 352, 353, 5, 189, 95, 2, 353,
	354, 5, 201, 101, 2, 354, 355, 5, 191, 96, 2, 355, 56, 3, 2, 2, 2, 356,
	357, 5, 175, 88, 2, 357, 358, 5, 161, 81, 2, 358, 359, 5, 203, 102, 2,
	359, 360, 5, 177, 89, 2, 360, 361, 5, 187, 94, 2, 361, 362, 5, 173, 87,
	2, 362, 58, 3, 2, 2, 2, 363, 364, 5, 177, 89, 2, 364, 365, 5, 173, 87,
	2, 365, 366, 5, 187, 94, 2, 366, 367, 5, 189, 95, 2, 367, 368, 5, 195,
	98, 2, 368, 369, 5, 177, 89, 2, 369, 370, 5, 187, 94, 2, 370, 371, 5, 173,
	87, 2, 371, 60, 3, 2, 2, 2, 372, 373, 5, 177, 89, 2, 373, 374, 5, 187,
	94, 2, 374, 62, 3, 2, 2, 2, 375, 376, 5, 177, 89, 2, 376, 377, 5, 187,
	94, 2, 377, 378, 5, 197, 99, 2, 378, 379, 5, 169, 85, 2, 379, 380, 5, 195,
	98, 2, 380, 381, 5, 199, 100, 2, 381, 64, 3, 2, 2, 2, 382, 383, 5, 177,
	89, 2, 383, 384, 5, 187, 94, 2, 384, 385, 5, 199, 100, 2, 385, 386, 5,
	189, 95, 2, 386, 66, 3, 2, 2, 2, 387, 388, 5, 179, 90, 2, 388, 389, 5,
	189, 95, 2, 389, 390, 5, 177, 89, 2, 390, 391, 5, 187, 94, 2, 391, 68,
	3, 2, 2, 2, 392, 393, 5, 183, 92, 2, 393, 394, 5, 161, 81, 2, 394, 395,
	5, 199, 100, 2, 395, 396, 5, 183, 92, 2, 396, 397, 5, 187, 94, 2, 397,
	398, 5, 173, 87, 2, 398, 70, 3, 2, 2, 2, 399, 400, 5, 183, 92, 2, 400,
	401, 5, 169, 85, 2, 401, 402, 5, 171, 86, 2, 402, 403, 5, 199, 100, 2,
	403, 72, 3, 2, 2, 2, 404, 405, 5, 183, 92, 2, 405, 406, 5, 177, 89, 2,
	406, 407, 5, 181, 91, 2, 407, 408, 5, 169, 85, 2, 408, 74, 3, 2, 2, 2,
	409, 410, 5, 183, 92, 2, 410, 411, 5, 177, 89, 2, 411, 412, 5, 185, 93,
	2, 412, 413, 5, 177, 89, 2, 413, 414, 5, 199, 100, 2, 414, 76, 3, 2, 2,
	2, 415, 416, 5, 185, 93, 2, 416, 417, 5, 161, 81, 2, 417, 418, 5, 199,
	100, 2, 418, 419, 5, 165, 83, 2, 419, 420, 5, 175, 88, 2, 420, 421, 5,
	169, 85, 2, 421, 422, 5, 197, 99, 2, 422, 78, 3, 2, 2, 2, 423, 424, 5,
	185, 93, 2, 424, 425, 5, 161, 81, 2, 425, 426, 5, 207, 104, 2, 426, 427,
	5, 177, 89, 2, 427, 428, 5, 185, 93, 2, 428, 429, 5, 201, 101, 2, 429,
	430, 5, 185, 93, 2, 430, 80, 3, 2, 2, 2, 431, 432, 5, 185, 93, 2, 432,
	433, 5, 177, 89, 2, 433, 434, 5, 187, 94, 2, 434, 435, 5, 177, 89, 2, 435,
	436, 5, 185, 93, 2, 436, 437, 5, 201, 101, 2, 437, 438, 5, 185, 93, 2,
	438, 82, 3, 2, 2, 2, 439, 440, 5, 187, 94, 2, 440, 441, 5, 189, 95, 2,
	441, 442, 5, 199, 100, 2, 442, 84, 3, 2, 2, 2, 443, 444, 5, 169, 85, 2,
	444, 445, 5, 193, 97, 2, 445, 446, 5, 201, 101, 2, 446, 447, 5, 161, 81,
	2, 447, 448, 5, 183, 92, 2, 448, 86, 3, 2, 2, 2, 449, 450, 5, 189, 95,
	2, 450, 451, 5, 171, 86, 2, 451, 88, 3, 2, 2, 2, 452, 453, 5, 189, 95,
	2, 453, 454, 5, 171, 86, 2, 454, 455, 5, 171, 86, 2, 455, 456, 5, 197,
	99, 2, 456, 457, 5, 169, 85, 2, 457, 458, 5, 199, 100, 2, 458, 90, 3, 2,
	2, 2, 459, 460, 5, 189, 95, 2, 460, 461, 5, 187, 94, 2, 461, 92, 3, 2,
	2, 2, 462, 463, 5, 189, 95, 2, 463, 464, 5, 195, 98, 2, 464, 465, 5, 167,
	84, 2, 465, 466, 5, 169, 85, 2, 466, 467, 5, 195, 98, 2, 467, 94, 3, 2,
	2, 2, 468, 469, 5, 189, 95, 2, 469, 470, 5, 201, 101, 2, 470, 471, 5, 199,
	100, 2, 471, 472, 5, 169, 85, 2, 472, 473, 5, 195, 98, 2, 473, 96, 3, 2,
	2, 2, 474, 475, 5, 195, 98, 2, 475, 476, 5, 169, 85, 2, 476, 477, 5, 165,
	83, 2, 477, 478, 5, 199, 100, 2, 478, 479, 5, 161, 81, 2, 479, 480, 5,
	187, 94, 2, 480, 481, 5, 173, 87, 2, 481, 482, 5, 183, 92, 2, 482, 483,
	5, 169, 85, 2, 483, 98, 3, 2, 2, 2, 484, 485, 5, 195, 98, 2, 485, 486,
	5, 169, 85, 2, 486, 487, 5, 187, 94, 2, 487, 488, 5, 161, 81, 2, 488, 489,
	5, 185, 93, 2, 489, 490, 5, 169, 85, 2, 490, 100, 3, 2, 2, 2, 491, 492,
	5, 197, 99, 2, 492, 493, 5, 199, 100, 2, 493, 494, 7, 97, 2, 2, 494, 495,
	5, 167, 84, 2, 495, 496, 5, 177, 89, 2, 496, 497, 5, 197, 99, 2, 497, 498,
	5, 199, 100, 2, 498, 499, 5, 161, 81, 2, 499, 500, 5, 187, 94, 2, 500,
	501, 5, 165, 83, 2, 501, 502, 5, 169, 85, 2, 502, 102, 3, 2, 2, 2, 503,
	504, 5, 197, 99, 2, 504, 505, 5, 169, 85, 2, 505, 506, 5, 183, 92, 2, 506,
	507, 5, 169, 85, 2, 507, 508, 5, 165, 83, 2, 508, 509, 5, 199, 100, 2,
	509, 104, 3, 2, 2, 2, 510, 511, 5, 197, 99, 2, 511, 512, 5, 199, 100, 2,
	512, 513, 7, 97, 2, 2, 513, 514, 5, 177, 89, 2, 514, 515, 5, 187, 94, 2,
	515, 516, 5, 199, 100, 2, 516, 517, 5, 169, 85, 2, 517, 518, 5, 195, 98,
	2, 518, 519, 5, 197, 99, 2, 519, 520, 5, 169, 85, 2, 520, 521, 5, 165,
	83, 2, 521, 522, 5, 199, 100, 2, 522, 523, 5, 197, 99, 2, 523, 106, 3,
	2, 2, 2, 524, 525, 5, 197, 99, 2, 525, 526, 5, 201, 101, 2, 526, 527, 5,
	185, 93, 2, 527, 108, 3, 2, 2, 2, 528, 529, 5, 197, 99, 2, 529, 530, 5,
	169, 85, 2, 530, 531, 5, 199, 100, 2, 531, 110, 3, 2, 2, 2, 532, 533, 5,
	197, 99, 2, 533, 534, 5, 175, 88, 2, 534, 535, 5, 189, 95, 2, 535, 536,
	5, 205, 103, 2, 536, 112, 3, 2, 2, 2, 537, 538, 5, 197, 99, 2, 538, 539,
	5, 199, 100, 2, 539, 540, 5, 161, 81, 2, 540, 541, 5, 195, 98, 2, 541,
	542, 5, 199, 100, 2, 542, 543, 5, 197, 99, 2, 543, 114, 3, 2, 2, 2, 544,
	545, 5, 199, 100, 2, 545, 546, 5, 161, 81, 2, 546, 547, 5, 163, 82, 2,
	547, 548, 5, 183, 92, 2, 548, 549, 5, 169, 85, 2, 549, 116, 3, 2, 2, 2,
	550, 551, 5, 199, 100, 2, 551, 552, 5, 161, 81, 2, 552, 553, 5, 163, 82,
	2, 553, 554, 5, 183, 92, 2, 554, 555, 5, 169, 85, 2, 555, 556, 5, 197,
	99, 2, 556, 118, 3, 2, 2, 2, 557, 558, 5, 199, 100, 2, 558, 559, 5, 189,
	95, 2, 559, 120, 3, 2, 2, 2, 560, 561, 5, 201, 101, 2, 561, 562, 5, 191,
	96, 2, 562, 563, 5, 167, 84, 2, 563, 564, 5, 161, 81, 2, 564, 565, 5, 199,
	100, 2, 565, 566, 5, 169, 85, 2, 566, 122, 3, 2, 2, 2, 567, 568, 5, 203,
	102, 2, 568, 569, 5, 161, 81, 2, 569, 570, 5, 183, 92, 2, 570, 571, 5,
	201, 101, 2, 571, 572, 5, 169, 85, 2, 572, 573, 5, 197, 99, 2, 573, 124,
	3, 2, 2, 2, 574, 575, 5, 203, 102, 2, 575, 576, 5, 177, 89, 2, 576, 577,
	5, 169, 85, 2, 577, 578, 5, 205, 103, 2, 578, 126, 3, 2, 2, 2, 579, 580,
	5, 205, 103, 2, 580, 581, 5, 175, 88, 2, 581, 582, 5, 169, 85, 2, 582,
	583, 5, 195, 98, 2, 583, 584, 5, 169, 85, 2, 584, 128, 3, 2, 2, 2, 585,
	586, 5, 205, 103, 2, 586, 587, 5, 177, 89, 2, 587, 588, 5, 199, 100, 2,
	588, 589, 5, 175, 88, 2, 589, 130, 3, 2, 2, 2, 590, 591, 7, 62, 2, 2, 591,
	592, 7, 63, 2, 2, 592, 132, 3, 2, 2, 2, 593, 594, 7, 64, 2, 2, 594, 595,
	7, 63, 2, 2, 595, 134, 3, 2, 2, 2, 596, 597, 7, 64, 2, 2, 597, 136, 3,
	2, 2, 2, 598, 599, 7, 63, 2, 2, 599, 138, 3, 2, 2, 2, 600, 601, 7, 62,
	2, 2, 601, 140, 3, 2, 2, 2, 602, 603, 7, 42, 2, 2, 603, 142, 3, 2, 2, 2,
	604, 605, 7, 43, 2, 2, 605, 144, 3, 2, 2, 2, 606, 608, 5, 159, 80, 2, 607,
	606, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609, 607, 3, 2, 2, 2, 609, 610,
	3, 2, 2, 2, 610, 618, 3, 2, 2, 2, 611, 615, 7, 48, 2, 2, 612, 614, 5, 159,
	80, 2, 613, 612, 3, 2, 2, 2, 614, 617, 3, 2, 2, 2, 615, 613, 3, 2, 2, 2,
	615, 616, 3, 2, 2, 2, 616, 619, 3, 2, 2, 2, 617, 615, 3, 2, 2, 2, 618,
	611, 3, 2, 2, 2, 618, 619, 3, 2, 2, 2, 619, 623, 3, 2, 2, 2, 620, 621,
	9, 2, 2, 2, 621, 623, 5, 145, 73, 2, 622, 607, 3, 2, 2, 2, 622, 620, 3,
	2, 2, 2, 623, 146, 3, 2, 2, 2, 624, 627, 5, 149, 75, 2, 625, 627, 5, 151,
	76, 2, 626, 624, 3, 2, 2, 2, 626, 625, 3, 2, 2, 2, 627, 148, 3, 2, 2, 2,
	628, 630, 9, 3, 2, 2, 629, 628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2, 631,
	629, 3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 150, 3, 2, 2, 2, 633, 639,
	7, 41, 2, 2, 634, 638, 10, 4, 2, 2, 635, 636, 7, 41, 2, 2, 636, 638, 7,
	41, 2, 2, 637, 634, 3, 2, 2, 2, 637, 635, 3, 2, 2, 2, 638, 641, 3, 2, 2,
	2, 639, 637, 3, 2, 2, 2, 639, 640, 3, 2, 2, 2, 640, 642, 3, 2, 2, 2, 641,
	639, 3, 2, 2, 2, 642, 643, 7, 41, 2, 2, 643, 152, 3, 2, 2, 2, 644, 645,
	7, 47, 2, 2, 645, 646, 7, 47, 2, 2, 646, 650, 3, 2, 2, 2, 647, 649, 10,
	5, 2, 2, 648, 647, 3, 2, 2, 2, 649, 652, 3, 2, 2, 2, 650, 648, 3, 2, 2,
	2, 650, 651, 3, 2, 2, 2, 651, 653, 3, 2, 2, 2, 652, 650, 3, 2, 2, 2, 653,
	654, 8, 77, 2, 2, 654, 154, 3, 2, 2, 2, 655, 656, 7, 49, 2, 2, 656, 657,
	7, 44, 2, 2, 657, 661, 3, 2, 2, 2, 658, 660, 11, 2, 2, 2, 659, 658, 3,
	2, 2, 2, 660, 663, 3, 2, 2, 2, 661, 662, 3, 2, 2, 2, 661, 659, 3, 2, 2,
	2, 662, 667, 3, 2, 2, 2, 663, 661, 3, 2, 2, 2, 664, 665, 7, 44, 2, 2, 665,
	668, 7, 49, 2, 2, 666, 668, 7, 2, 2, 3, 667, 664, 3, 2, 2, 2, 667, 666,
	3, 2, 2, 2, 668, 669, 3, 2, 2, 2, 669, 670, 8, 78, 2, 2, 670, 156, 3, 2,
	2, 2, 671, 672, 9, 6, 2, 2, 672, 673, 3, 2, 2, 2, 673, 674, 8, 79, 2, 2,
	674, 158, 3, 2, 2, 2, 675, 676, 9, 7, 2, 2, 676, 160, 3, 2, 2, 2, 677,
	678, 9, 8, 2, 2, 678, 162, 3, 2, 2, 2, 679, 680, 9, 9, 2, 2, 680, 164,
	3, 2, 2, 2, 681, 682, 9, 10, 2, 2, 682, 166, 3, 2, 2, 2, 683, 684, 9, 11,
	2, 2, 684, 168, 3, 2, 2, 2, 685, 686, 9, 12, 2, 2, 686, 170, 3, 2, 2, 2,
	687, 688, 9, 13, 2, 2, 688, 172, 3, 2, 2, 2, 689, 690, 9, 14, 2, 2, 690,
	174, 3, 2, 2, 2, 691, 692, 9, 15, 2, 2, 692, 176, 3, 2, 2, 2, 693, 694,
	9, 16, 2, 2, 694, 178, 3, 2, 2, 2, 695, 696, 9, 17, 2, 2, 696, 180, 3,
	2, 2, 2, 697, 698, 9, 18, 2, 2, 698, 182, 3, 2, 2, 2, 699, 700, 9, 19,
	2, 2, 700, 184, 3, 2, 2, 2, 701, 702, 9, 20, 2, 2, 702, 186, 3, 2, 2, 2,
	703, 704, 9, 21, 2, 2, 704, 188, 3, 2, 2, 2, 705, 706, 9, 22, 2, 2, 706,
	190, 3, 2, 2, 2, 707, 708, 9, 23, 2, 2, 708, 192, 3, 2, 2, 2, 709, 710,
	9, 24, 2, 2, 710, 194, 3, 2, 2, 2, 711, 712, 9, 25, 2, 2, 712, 196, 3,
	2, 2, 2, 713, 714, 9, 26, 2, 2, 714, 198, 3, 2, 2, 2, 715, 716, 9, 27,
	2, 2, 716, 200, 3, 2, 2, 2, 717, 718, 9, 28, 2, 2, 718, 202, 3, 2, 2, 2,
	719, 720, 9, 29, 2, 2, 720, 204, 3, 2, 2, 2, 721, 722, 9, 30, 2, 2, 722,
	206, 3, 2, 2, 2, 723, 724, 9, 31, 2, 2, 724, 208, 3, 2, 2, 2, 725, 726,
	9, 32, 2, 2, 726, 210, 3, 2, 2, 2, 727, 728, 9, 33, 2, 2, 728, 212, 3,
	2, 2, 2, 15, 2, 609, 615, 618, 622, 626, 629, 631, 637, 639, 650, 661,
	667, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'*'", "','", "'.'", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "'<='", "'>='",
	"'>'", "'='", "'<'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "K_ALTER", "K_AND", "K_OR", "K_AS", "K_ASC", "K_AVERAGE",
	"K_BY", "K_BETWEEN", "K_CASE", "K_CIRCLE", "K_COLUMN", "K_CONTAINS", "K_COUNT",
	"K_CREATE", "K_DELETE", "K_DESC", "K_DESCRIBE", "K_DOES", "K_CONTAIN",
	"K_DROP", "K_ENDS", "K_FROM", "K_GROUP", "K_HAVING", "K_IGNORING", "K_IN",
	"K_INSERT", "K_INTO", "K_JOIN", "K_LATLNG", "K_LEFT", "K_LIKE", "K_LIMIT",
	"K_MATCHES", "K_MAXIMUM", "K_MINIMUM", "K_NOT", "K_EQUAL", "K_OF", "K_OFFSET",
	"K_ON", "K_ORDER", "K_OUTER", "K_RECTANGLE", "K_RENAME", "K_ST_DISTANCE",
	"K_SELECT", "K_ST_INTERSECTS", "K_SUM", "K_SET", "K_SHOW", "K_STARTS",
	"K_TABLE", "K_TABLES", "K_TO", "K_UPDATE", "K_VALUES", "K_VIEW", "K_WHERE",
	"K_WITH", "LT_EQ", "GT_EQ", "GT", "EQ", "LT", "LPAR", "RPAR", "NUMERIC_LITERAL",
	"STRING_LITERAL", "STRING", "QUOTED_STRING", "SINGLELINE_COMMENT", "MULTILINE_COMMENT",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "K_ALTER", "K_AND", "K_OR", "K_AS", "K_ASC",
	"K_AVERAGE", "K_BY", "K_BETWEEN", "K_CASE", "K_CIRCLE", "K_COLUMN", "K_CONTAINS",
	"K_COUNT", "K_CREATE", "K_DELETE", "K_DESC", "K_DESCRIBE", "K_DOES", "K_CONTAIN",
	"K_DROP", "K_ENDS", "K_FROM", "K_GROUP", "K_HAVING", "K_IGNORING", "K_IN",
	"K_INSERT", "K_INTO", "K_JOIN", "K_LATLNG", "K_LEFT", "K_LIKE", "K_LIMIT",
	"K_MATCHES", "K_MAXIMUM", "K_MINIMUM", "K_NOT", "K_EQUAL", "K_OF", "K_OFFSET",
	"K_ON", "K_ORDER", "K_OUTER", "K_RECTANGLE", "K_RENAME", "K_ST_DISTANCE",
	"K_SELECT", "K_ST_INTERSECTS", "K_SUM", "K_SET", "K_SHOW", "K_STARTS",
	"K_TABLE", "K_TABLES", "K_TO", "K_UPDATE", "K_VALUES", "K_VIEW", "K_WHERE",
	"K_WITH", "LT_EQ", "GT_EQ", "GT", "EQ", "LT", "LPAR", "RPAR", "NUMERIC_LITERAL",
	"STRING_LITERAL", "STRING", "QUOTED_STRING", "SINGLELINE_COMMENT", "MULTILINE_COMMENT",
	"WHITESPACE", "DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
	"Z",
}

type FusionTablesSqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewFusionTablesSqlLexer(input antlr.CharStream) *FusionTablesSqlLexer {

	l := new(FusionTablesSqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FusionTablesSql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FusionTablesSqlLexer tokens.
const (
	FusionTablesSqlLexerT__0               = 1
	FusionTablesSqlLexerT__1               = 2
	FusionTablesSqlLexerT__2               = 3
	FusionTablesSqlLexerT__3               = 4
	FusionTablesSqlLexerK_ALTER            = 5
	FusionTablesSqlLexerK_AND              = 6
	FusionTablesSqlLexerK_OR               = 7
	FusionTablesSqlLexerK_AS               = 8
	FusionTablesSqlLexerK_ASC              = 9
	FusionTablesSqlLexerK_AVERAGE          = 10
	FusionTablesSqlLexerK_BY               = 11
	FusionTablesSqlLexerK_BETWEEN          = 12
	FusionTablesSqlLexerK_CASE             = 13
	FusionTablesSqlLexerK_CIRCLE           = 14
	FusionTablesSqlLexerK_COLUMN           = 15
	FusionTablesSqlLexerK_CONTAINS         = 16
	FusionTablesSqlLexerK_COUNT            = 17
	FusionTablesSqlLexerK_CREATE           = 18
	FusionTablesSqlLexerK_DELETE           = 19
	FusionTablesSqlLexerK_DESC             = 20
	FusionTablesSqlLexerK_DESCRIBE         = 21
	FusionTablesSqlLexerK_DOES             = 22
	FusionTablesSqlLexerK_CONTAIN          = 23
	FusionTablesSqlLexerK_DROP             = 24
	FusionTablesSqlLexerK_ENDS             = 25
	FusionTablesSqlLexerK_FROM             = 26
	FusionTablesSqlLexerK_GROUP            = 27
	FusionTablesSqlLexerK_HAVING           = 28
	FusionTablesSqlLexerK_IGNORING         = 29
	FusionTablesSqlLexerK_IN               = 30
	FusionTablesSqlLexerK_INSERT           = 31
	FusionTablesSqlLexerK_INTO             = 32
	FusionTablesSqlLexerK_JOIN             = 33
	FusionTablesSqlLexerK_LATLNG           = 34
	FusionTablesSqlLexerK_LEFT             = 35
	FusionTablesSqlLexerK_LIKE             = 36
	FusionTablesSqlLexerK_LIMIT            = 37
	FusionTablesSqlLexerK_MATCHES          = 38
	FusionTablesSqlLexerK_MAXIMUM          = 39
	FusionTablesSqlLexerK_MINIMUM          = 40
	FusionTablesSqlLexerK_NOT              = 41
	FusionTablesSqlLexerK_EQUAL            = 42
	FusionTablesSqlLexerK_OF               = 43
	FusionTablesSqlLexerK_OFFSET           = 44
	FusionTablesSqlLexerK_ON               = 45
	FusionTablesSqlLexerK_ORDER            = 46
	FusionTablesSqlLexerK_OUTER            = 47
	FusionTablesSqlLexerK_RECTANGLE        = 48
	FusionTablesSqlLexerK_RENAME           = 49
	FusionTablesSqlLexerK_ST_DISTANCE      = 50
	FusionTablesSqlLexerK_SELECT           = 51
	FusionTablesSqlLexerK_ST_INTERSECTS    = 52
	FusionTablesSqlLexerK_SUM              = 53
	FusionTablesSqlLexerK_SET              = 54
	FusionTablesSqlLexerK_SHOW             = 55
	FusionTablesSqlLexerK_STARTS           = 56
	FusionTablesSqlLexerK_TABLE            = 57
	FusionTablesSqlLexerK_TABLES           = 58
	FusionTablesSqlLexerK_TO               = 59
	FusionTablesSqlLexerK_UPDATE           = 60
	FusionTablesSqlLexerK_VALUES           = 61
	FusionTablesSqlLexerK_VIEW             = 62
	FusionTablesSqlLexerK_WHERE            = 63
	FusionTablesSqlLexerK_WITH             = 64
	FusionTablesSqlLexerLT_EQ              = 65
	FusionTablesSqlLexerGT_EQ              = 66
	FusionTablesSqlLexerGT                 = 67
	FusionTablesSqlLexerEQ                 = 68
	FusionTablesSqlLexerLT                 = 69
	FusionTablesSqlLexerLPAR               = 70
	FusionTablesSqlLexerRPAR               = 71
	FusionTablesSqlLexerNUMERIC_LITERAL    = 72
	FusionTablesSqlLexerSTRING_LITERAL     = 73
	FusionTablesSqlLexerSTRING             = 74
	FusionTablesSqlLexerQUOTED_STRING      = 75
	FusionTablesSqlLexerSINGLELINE_COMMENT = 76
	FusionTablesSqlLexerMULTILINE_COMMENT  = 77
	FusionTablesSqlLexerWHITESPACE         = 78
)
