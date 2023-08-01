package irsdk

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strings"
)

func byte4ToInt(in []byte) int {
	return int(binary.LittleEndian.Uint32(in))
}
func byte4ToFloat(in []byte) float32 {
	bits := binary.LittleEndian.Uint32(in)
	return math.Float32frombits(bits)
}
func byte4sToFloats(in []byte) []float32 {
	if len(in)%4 != 0 {
		log.Fatal("expect byte array of float32 (4 bytes)")
	}
	count := len(in) / 4
	result := make([]float32, count)
	for i := 0; i < count; i++ {
		bits := binary.LittleEndian.Uint32(in[i*4 : i*4+4])
		result[i] = math.Float32frombits(bits)
	}
	return result
}
func byte8ToFloat(in []byte) float64 {
	bits := binary.LittleEndian.Uint64(in)
	return math.Float64frombits(bits)
}
func byte4toBitField(in []byte) string {
	return fmt.Sprintf("0x%x", int(binary.LittleEndian.Uint32(in)))
}
func bytesToString(in []byte) string {
	return strings.TrimRight(string(in), "\x00")
}
