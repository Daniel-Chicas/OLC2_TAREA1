// Code generated from C:/Users/Daniel Chicas/Desktop/PRIMER_SEMESTRE_2022/COMPI2/LABORATORIO/TAREAS_LAB/OLC2_TAREA1/Backend/Gramatica\Gramatica.g4 by ANTLR 4.9.2. DO NOT EDIT.

package Gramatica

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 48, 422,
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
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 7, 42, 303, 10, 42, 12, 42, 14, 42, 306, 11, 42, 3, 42,
	3, 42, 3, 43, 6, 43, 311, 10, 43, 13, 43, 14, 43, 312, 3, 44, 6, 44, 316,
	10, 44, 13, 44, 14, 44, 317, 3, 44, 3, 44, 7, 44, 322, 10, 44, 12, 44,
	14, 44, 325, 11, 44, 3, 44, 3, 44, 6, 44, 329, 10, 44, 13, 44, 14, 44,
	330, 5, 44, 333, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 6, 47, 347, 10, 47, 13, 47, 14, 47,
	348, 3, 47, 3, 47, 7, 47, 353, 10, 47, 12, 47, 14, 47, 356, 11, 47, 3,
	48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52,
	3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3,
	55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60,
	3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3,
	66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71,
	3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 6, 74, 417, 10, 74, 13, 74, 14,
	74, 418, 3, 74, 3, 74, 2, 2, 75, 3, 2, 5, 2, 7, 2, 9, 2, 11, 2, 13, 2,
	15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 33, 2, 35,
	2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 2,
	57, 3, 59, 4, 61, 5, 63, 6, 65, 7, 67, 8, 69, 9, 71, 10, 73, 11, 75, 12,
	77, 13, 79, 14, 81, 15, 83, 16, 85, 17, 87, 18, 89, 19, 91, 20, 93, 21,
	95, 22, 97, 23, 99, 24, 101, 25, 103, 26, 105, 27, 107, 28, 109, 29, 111,
	30, 113, 31, 115, 32, 117, 33, 119, 34, 121, 35, 123, 36, 125, 37, 127,
	38, 129, 39, 131, 40, 133, 41, 135, 42, 137, 43, 139, 44, 141, 45, 143,
	46, 145, 47, 147, 48, 3, 2, 34, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100,
	100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103,
	103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106,
	106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109,
	109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112,
	112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115,
	115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118,
	118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121,
	121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124,
	124, 9, 2, 34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 4, 2,
	36, 36, 94, 94, 3, 2, 50, 59, 14, 2, 67, 92, 99, 124, 163, 163, 171, 171,
	175, 175, 181, 181, 188, 188, 197, 197, 355, 355, 8222, 8222, 8242, 8242,
	65535, 65535, 15, 2, 50, 59, 67, 92, 99, 124, 163, 163, 171, 171, 175,
	175, 181, 181, 188, 188, 197, 197, 355, 355, 8222, 8222, 8242, 8242, 65535,
	65535, 5, 2, 11, 12, 15, 15, 34, 34, 2, 406, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119,
	3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2,
	2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3,
	2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2,
	141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2,
	2, 2, 3, 149, 3, 2, 2, 2, 5, 151, 3, 2, 2, 2, 7, 153, 3, 2, 2, 2, 9, 155,
	3, 2, 2, 2, 11, 157, 3, 2, 2, 2, 13, 159, 3, 2, 2, 2, 15, 161, 3, 2, 2,
	2, 17, 163, 3, 2, 2, 2, 19, 165, 3, 2, 2, 2, 21, 167, 3, 2, 2, 2, 23, 169,
	3, 2, 2, 2, 25, 171, 3, 2, 2, 2, 27, 173, 3, 2, 2, 2, 29, 175, 3, 2, 2,
	2, 31, 177, 3, 2, 2, 2, 33, 179, 3, 2, 2, 2, 35, 181, 3, 2, 2, 2, 37, 183,
	3, 2, 2, 2, 39, 185, 3, 2, 2, 2, 41, 187, 3, 2, 2, 2, 43, 189, 3, 2, 2,
	2, 45, 191, 3, 2, 2, 2, 47, 193, 3, 2, 2, 2, 49, 195, 3, 2, 2, 2, 51, 197,
	3, 2, 2, 2, 53, 199, 3, 2, 2, 2, 55, 201, 3, 2, 2, 2, 57, 204, 3, 2, 2,
	2, 59, 207, 3, 2, 2, 2, 61, 212, 3, 2, 2, 2, 63, 230, 3, 2, 2, 2, 65, 235,
	3, 2, 2, 2, 67, 243, 3, 2, 2, 2, 69, 249, 3, 2, 2, 2, 71, 258, 3, 2, 2,
	2, 73, 265, 3, 2, 2, 2, 75, 269, 3, 2, 2, 2, 77, 274, 3, 2, 2, 2, 79, 279,
	3, 2, 2, 2, 81, 289, 3, 2, 2, 2, 83, 296, 3, 2, 2, 2, 85, 310, 3, 2, 2,
	2, 87, 332, 3, 2, 2, 2, 89, 334, 3, 2, 2, 2, 91, 339, 3, 2, 2, 2, 93, 346,
	3, 2, 2, 2, 95, 357, 3, 2, 2, 2, 97, 360, 3, 2, 2, 2, 99, 363, 3, 2, 2,
	2, 101, 365, 3, 2, 2, 2, 103, 367, 3, 2, 2, 2, 105, 370, 3, 2, 2, 2, 107,
	373, 3, 2, 2, 2, 109, 376, 3, 2, 2, 2, 111, 379, 3, 2, 2, 2, 113, 381,
	3, 2, 2, 2, 115, 383, 3, 2, 2, 2, 117, 385, 3, 2, 2, 2, 119, 387, 3, 2,
	2, 2, 121, 389, 3, 2, 2, 2, 123, 391, 3, 2, 2, 2, 125, 393, 3, 2, 2, 2,
	127, 395, 3, 2, 2, 2, 129, 397, 3, 2, 2, 2, 131, 399, 3, 2, 2, 2, 133,
	401, 3, 2, 2, 2, 135, 403, 3, 2, 2, 2, 137, 405, 3, 2, 2, 2, 139, 407,
	3, 2, 2, 2, 141, 409, 3, 2, 2, 2, 143, 411, 3, 2, 2, 2, 145, 413, 3, 2,
	2, 2, 147, 416, 3, 2, 2, 2, 149, 150, 9, 2, 2, 2, 150, 4, 3, 2, 2, 2, 151,
	152, 9, 3, 2, 2, 152, 6, 3, 2, 2, 2, 153, 154, 9, 4, 2, 2, 154, 8, 3, 2,
	2, 2, 155, 156, 9, 5, 2, 2, 156, 10, 3, 2, 2, 2, 157, 158, 9, 6, 2, 2,
	158, 12, 3, 2, 2, 2, 159, 160, 9, 7, 2, 2, 160, 14, 3, 2, 2, 2, 161, 162,
	9, 8, 2, 2, 162, 16, 3, 2, 2, 2, 163, 164, 9, 9, 2, 2, 164, 18, 3, 2, 2,
	2, 165, 166, 9, 10, 2, 2, 166, 20, 3, 2, 2, 2, 167, 168, 9, 11, 2, 2, 168,
	22, 3, 2, 2, 2, 169, 170, 9, 12, 2, 2, 170, 24, 3, 2, 2, 2, 171, 172, 9,
	13, 2, 2, 172, 26, 3, 2, 2, 2, 173, 174, 9, 14, 2, 2, 174, 28, 3, 2, 2,
	2, 175, 176, 9, 15, 2, 2, 176, 30, 3, 2, 2, 2, 177, 178, 9, 16, 2, 2, 178,
	32, 3, 2, 2, 2, 179, 180, 9, 17, 2, 2, 180, 34, 3, 2, 2, 2, 181, 182, 9,
	18, 2, 2, 182, 36, 3, 2, 2, 2, 183, 184, 9, 19, 2, 2, 184, 38, 3, 2, 2,
	2, 185, 186, 9, 20, 2, 2, 186, 40, 3, 2, 2, 2, 187, 188, 9, 21, 2, 2, 188,
	42, 3, 2, 2, 2, 189, 190, 9, 22, 2, 2, 190, 44, 3, 2, 2, 2, 191, 192, 9,
	23, 2, 2, 192, 46, 3, 2, 2, 2, 193, 194, 9, 24, 2, 2, 194, 48, 3, 2, 2,
	2, 195, 196, 9, 25, 2, 2, 196, 50, 3, 2, 2, 2, 197, 198, 9, 26, 2, 2, 198,
	52, 3, 2, 2, 2, 199, 200, 9, 27, 2, 2, 200, 54, 3, 2, 2, 2, 201, 202, 7,
	94, 2, 2, 202, 203, 9, 28, 2, 2, 203, 56, 3, 2, 2, 2, 204, 205, 5, 19,
	10, 2, 205, 206, 5, 13, 7, 2, 206, 58, 3, 2, 2, 2, 207, 208, 5, 11, 6,
	2, 208, 209, 5, 25, 13, 2, 209, 210, 5, 39, 20, 2, 210, 211, 5, 11, 6,
	2, 211, 60, 3, 2, 2, 2, 212, 213, 5, 39, 20, 2, 213, 214, 5, 11, 6, 2,
	214, 215, 5, 29, 15, 2, 215, 216, 5, 41, 21, 2, 216, 217, 5, 11, 6, 2,
	217, 218, 5, 29, 15, 2, 218, 219, 5, 7, 4, 2, 219, 220, 5, 19, 10, 2, 220,
	221, 5, 3, 2, 2, 221, 222, 5, 127, 64, 2, 222, 223, 5, 7, 4, 2, 223, 224,
	5, 31, 16, 2, 224, 225, 5, 29, 15, 2, 225, 226, 5, 39, 20, 2, 226, 227,
	5, 31, 16, 2, 227, 228, 5, 25, 13, 2, 228, 229, 5, 3, 2, 2, 229, 62, 3,
	2, 2, 2, 230, 231, 5, 27, 14, 2, 231, 232, 5, 3, 2, 2, 232, 233, 5, 19,
	10, 2, 233, 234, 5, 29, 15, 2, 234, 64, 3, 2, 2, 2, 235, 236, 5, 33, 17,
	2, 236, 237, 5, 43, 22, 2, 237, 238, 5, 5, 3, 2, 238, 239, 5, 25, 13, 2,
	239, 240, 5, 19, 10, 2, 240, 241, 5, 7, 4, 2, 241, 242, 5, 31, 16, 2, 242,
	66, 3, 2, 2, 2, 243, 244, 5, 7, 4, 2, 244, 245, 5, 25, 13, 2, 245, 246,
	5, 3, 2, 2, 246, 247, 5, 39, 20, 2, 247, 248, 5, 39, 20, 2, 248, 68, 3,
	2, 2, 2, 249, 250, 5, 9, 5, 2, 250, 251, 5, 11, 6, 2, 251, 252, 5, 7, 4,
	2, 252, 253, 5, 25, 13, 2, 253, 254, 5, 3, 2, 2, 254, 255, 5, 37, 19, 2,
	255, 256, 5, 3, 2, 2, 256, 257, 5, 37, 19, 2, 257, 70, 3, 2, 2, 2, 258,
	259, 5, 39, 20, 2, 259, 260, 5, 41, 21, 2, 260, 261, 5, 37, 19, 2, 261,
	262, 5, 19, 10, 2, 262, 263, 5, 29, 15, 2, 263, 264, 5, 15, 8, 2, 264,
	72, 3, 2, 2, 2, 265, 266, 5, 19, 10, 2, 266, 267, 5, 29, 15, 2, 267, 268,
	5, 41, 21, 2, 268, 74, 3, 2, 2, 2, 269, 270, 5, 37, 19, 2, 270, 271, 5,
	11, 6, 2, 271, 272, 5, 3, 2, 2, 272, 273, 5, 25, 13, 2, 273, 76, 3, 2,
	2, 2, 274, 275, 5, 5, 3, 2, 275, 276, 5, 31, 16, 2, 276, 277, 5, 31, 16,
	2, 277, 278, 5, 25, 13, 2, 278, 78, 3, 2, 2, 2, 279, 280, 5, 33, 17, 2,
	280, 281, 5, 37, 19, 2, 281, 282, 5, 19, 10, 2, 282, 283, 5, 29, 15, 2,
	283, 284, 5, 7, 4, 2, 284, 285, 5, 19, 10, 2, 285, 286, 5, 33, 17, 2, 286,
	287, 5, 3, 2, 2, 287, 288, 5, 25, 13, 2, 288, 80, 3, 2, 2, 2, 289, 290,
	5, 27, 14, 2, 290, 291, 5, 11, 6, 2, 291, 292, 5, 41, 21, 2, 292, 293,
	5, 31, 16, 2, 293, 294, 5, 9, 5, 2, 294, 295, 5, 31, 16, 2, 295, 82, 3,
	2, 2, 2, 296, 304, 7, 36, 2, 2, 297, 298, 7, 94, 2, 2, 298, 303, 11, 2,
	2, 2, 299, 300, 7, 36, 2, 2, 300, 303, 7, 36, 2, 2, 301, 303, 10, 29, 2,
	2, 302, 297, 3, 2, 2, 2, 302, 299, 3, 2, 2, 2, 302, 301, 3, 2, 2, 2, 303,
	306, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307,
	3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 307, 308, 7, 36, 2, 2, 308, 84, 3, 2,
	2, 2, 309, 311, 9, 30, 2, 2, 310, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2,
	312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 86, 3, 2, 2, 2, 314, 316,
	9, 30, 2, 2, 315, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 315, 3, 2,
	2, 2, 317, 318, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 323, 5, 127, 64,
	2, 320, 322, 9, 30, 2, 2, 321, 320, 3, 2, 2, 2, 322, 325, 3, 2, 2, 2, 323,
	321, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 333, 3, 2, 2, 2, 325, 323,
	3, 2, 2, 2, 326, 328, 5, 127, 64, 2, 327, 329, 9, 30, 2, 2, 328, 327, 3,
	2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 330, 331, 3, 2, 2,
	2, 331, 333, 3, 2, 2, 2, 332, 315, 3, 2, 2, 2, 332, 326, 3, 2, 2, 2, 333,
	88, 3, 2, 2, 2, 334, 335, 5, 41, 21, 2, 335, 336, 5, 37, 19, 2, 336, 337,
	5, 43, 22, 2, 337, 338, 5, 11, 6, 2, 338, 90, 3, 2, 2, 2, 339, 340, 5,
	13, 7, 2, 340, 341, 5, 3, 2, 2, 341, 342, 5, 25, 13, 2, 342, 343, 5, 39,
	20, 2, 343, 344, 5, 11, 6, 2, 344, 92, 3, 2, 2, 2, 345, 347, 9, 31, 2,
	2, 346, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348,
	349, 3, 2, 2, 2, 349, 354, 3, 2, 2, 2, 350, 353, 5, 125, 63, 2, 351, 353,
	9, 32, 2, 2, 352, 350, 3, 2, 2, 2, 352, 351, 3, 2, 2, 2, 353, 356, 3, 2,
	2, 2, 354, 352, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 94, 3, 2, 2, 2,
	356, 354, 3, 2, 2, 2, 357, 358, 7, 62, 2, 2, 358, 359, 7, 63, 2, 2, 359,
	96, 3, 2, 2, 2, 360, 361, 7, 64, 2, 2, 361, 362, 7, 63, 2, 2, 362, 98,
	3, 2, 2, 2, 363, 364, 7, 62, 2, 2, 364, 100, 3, 2, 2, 2, 365, 366, 7, 64,
	2, 2, 366, 102, 3, 2, 2, 2, 367, 368, 7, 63, 2, 2, 368, 369, 7, 63, 2,
	2, 369, 104, 3, 2, 2, 2, 370, 371, 7, 35, 2, 2, 371, 372, 7, 63, 2, 2,
	372, 106, 3, 2, 2, 2, 373, 374, 7, 40, 2, 2, 374, 375, 7, 40, 2, 2, 375,
	108, 3, 2, 2, 2, 376, 377, 7, 126, 2, 2, 377, 378, 7, 126, 2, 2, 378, 110,
	3, 2, 2, 2, 379, 380, 7, 35, 2, 2, 380, 112, 3, 2, 2, 2, 381, 382, 7, 39,
	2, 2, 382, 114, 3, 2, 2, 2, 383, 384, 7, 44, 2, 2, 384, 116, 3, 2, 2, 2,
	385, 386, 7, 49, 2, 2, 386, 118, 3, 2, 2, 2, 387, 388, 7, 45, 2, 2, 388,
	120, 3, 2, 2, 2, 389, 390, 7, 47, 2, 2, 390, 122, 3, 2, 2, 2, 391, 392,
	7, 63, 2, 2, 392, 124, 3, 2, 2, 2, 393, 394, 7, 97, 2, 2, 394, 126, 3,
	2, 2, 2, 395, 396, 7, 48, 2, 2, 396, 128, 3, 2, 2, 2, 397, 398, 7, 42,
	2, 2, 398, 130, 3, 2, 2, 2, 399, 400, 7, 43, 2, 2, 400, 132, 3, 2, 2, 2,
	401, 402, 7, 125, 2, 2, 402, 134, 3, 2, 2, 2, 403, 404, 7, 127, 2, 2, 404,
	136, 3, 2, 2, 2, 405, 406, 7, 93, 2, 2, 406, 138, 3, 2, 2, 2, 407, 408,
	7, 95, 2, 2, 408, 140, 3, 2, 2, 2, 409, 410, 7, 60, 2, 2, 410, 142, 3,
	2, 2, 2, 411, 412, 7, 61, 2, 2, 412, 144, 3, 2, 2, 2, 413, 414, 7, 46,
	2, 2, 414, 146, 3, 2, 2, 2, 415, 417, 9, 33, 2, 2, 416, 415, 3, 2, 2, 2,
	417, 418, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419,
	420, 3, 2, 2, 2, 420, 421, 8, 74, 2, 2, 421, 148, 3, 2, 2, 2, 14, 2, 302,
	304, 312, 317, 323, 330, 332, 348, 352, 354, 418, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "'<='", "'>='", "'<'", "'>'", "'=='", "'!='", "'&&'", "'||'", "'!'",
	"'%'", "'*'", "'/'", "'+'", "'-'", "'='", "'_'", "'.'", "'('", "')'", "'{'",
	"'}'", "'['", "']'", "':'", "';'", "','",
}

var lexerSymbolicNames = []string{
	"", "IF", "ELSE", "IMPRIMIR", "MAIN", "PUBLIC", "CLASS", "DECLARAR", "STRING",
	"INT", "REAL", "BOOLEAN", "PRINCIPAL", "METODO", "CADENA", "ENTERO", "DECIMAL",
	"TRUE", "FALSE", "ID", "MENI", "MAYI", "MEN", "MAY", "IGUALI", "DIFERENCIA",
	"AND", "OR", "NOT", "MOD", "POR", "DIVIDIR", "MAS", "MENOS", "IGUAL", "GUIONB",
	"PUNTO", "PARA", "PARC", "LLABRE", "LLACIE", "CABRE", "CCIER", "DPUNTOS",
	"PCOMA", "COMA", "ESPACIOS",
}

var lexerRuleNames = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "ESC_SEQ", "IF",
	"ELSE", "IMPRIMIR", "MAIN", "PUBLIC", "CLASS", "DECLARAR", "STRING", "INT",
	"REAL", "BOOLEAN", "PRINCIPAL", "METODO", "CADENA", "ENTERO", "DECIMAL",
	"TRUE", "FALSE", "ID", "MENI", "MAYI", "MEN", "MAY", "IGUALI", "DIFERENCIA",
	"AND", "OR", "NOT", "MOD", "POR", "DIVIDIR", "MAS", "MENOS", "IGUAL", "GUIONB",
	"PUNTO", "PARA", "PARC", "LLABRE", "LLACIE", "CABRE", "CCIER", "DPUNTOS",
	"PCOMA", "COMA", "ESPACIOS",
}

type GramaticaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewGramaticaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *GramaticaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaLexer(input antlr.CharStream) *GramaticaLexer {
	l := new(GramaticaLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GramaticaLexer tokens.
const (
	GramaticaLexerIF         = 1
	GramaticaLexerELSE       = 2
	GramaticaLexerIMPRIMIR   = 3
	GramaticaLexerMAIN       = 4
	GramaticaLexerPUBLIC     = 5
	GramaticaLexerCLASS      = 6
	GramaticaLexerDECLARAR   = 7
	GramaticaLexerSTRING     = 8
	GramaticaLexerINT        = 9
	GramaticaLexerREAL       = 10
	GramaticaLexerBOOLEAN    = 11
	GramaticaLexerPRINCIPAL  = 12
	GramaticaLexerMETODO     = 13
	GramaticaLexerCADENA     = 14
	GramaticaLexerENTERO     = 15
	GramaticaLexerDECIMAL    = 16
	GramaticaLexerTRUE       = 17
	GramaticaLexerFALSE      = 18
	GramaticaLexerID         = 19
	GramaticaLexerMENI       = 20
	GramaticaLexerMAYI       = 21
	GramaticaLexerMEN        = 22
	GramaticaLexerMAY        = 23
	GramaticaLexerIGUALI     = 24
	GramaticaLexerDIFERENCIA = 25
	GramaticaLexerAND        = 26
	GramaticaLexerOR         = 27
	GramaticaLexerNOT        = 28
	GramaticaLexerMOD        = 29
	GramaticaLexerPOR        = 30
	GramaticaLexerDIVIDIR    = 31
	GramaticaLexerMAS        = 32
	GramaticaLexerMENOS      = 33
	GramaticaLexerIGUAL      = 34
	GramaticaLexerGUIONB     = 35
	GramaticaLexerPUNTO      = 36
	GramaticaLexerPARA       = 37
	GramaticaLexerPARC       = 38
	GramaticaLexerLLABRE     = 39
	GramaticaLexerLLACIE     = 40
	GramaticaLexerCABRE      = 41
	GramaticaLexerCCIER      = 42
	GramaticaLexerDPUNTOS    = 43
	GramaticaLexerPCOMA      = 44
	GramaticaLexerCOMA       = 45
	GramaticaLexerESPACIOS   = 46
)
