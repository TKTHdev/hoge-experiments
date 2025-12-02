#include "textflag.h"

// func DotAVX2(x, y []float32) float32
// x.ptr: 0(FP), x.len: 8(FP), x.cap: 16(FP)
// y.ptr: 24(FP), y.len: 32(FP), y.cap: 40(FP)
// ret: 48(FP)
TEXT ·DotAVX2(SB), NOSPLIT, $0-52
	// Offsets based on slice header offsets.
	// To check, use `GOARCH=amd64 go vet`
	MOVQ a_base+0(FP), AX
	MOVQ b_base+24(FP), BX
	MOVQ a_len+8(FP), DX
	XORQ R8, R8 // return sum
	// Zero Y0, which will store 8 packed 32-bit sums
	VPXOR Y0, Y0, Y0
// In blockloop, we calculate the dot product 16 at a time
blockloop:
	CMPQ DX, $16
	JB reduce
	// Sign-extend 16 bytes into 16 int16s
	VPMOVSXBW (AX), Y1
	VPMOVSXBW (BX), Y2
	// Multiply words vertically to form doubleword intermediates,
	// then add adjacent doublewords.
	VPMADDWD Y1, Y2, Y1
	// Add results to the running sum
	VPADDD Y0, Y1, Y0
	ADDQ $16, AX
	ADDQ $16, BX
	SUBQ $16, DX
	JMP blockloop
reduce:
	// X0 is the low bits of Y0.
	// Extract the high bits into X1, fold in half, add, repeat.
	VEXTRACTI128 $1, Y0, X1
	VPADDD X0, X1, X0
	VPSRLDQ $8, X0, X1
	VPADDD X0, X1, X0
	VPSRLDQ $4, X0, X1
	VPADDD X0, X1, X0
	// Store the reduced sum
	VMOVD X0, R8
end:
	MOVL R8, ret+48(FP)
	VZEROALL
	RET
    