	.text
	.intel_syntax noprefix
	.file    "_FMA.c"
	.globl    MultiplyBytes                   # -- Begin function MultiplyBytes
	.p2align    4, 0x90
	.type    MultiplyBytes,@function
MultiplyBytes:                          # @MultiplyBytes
# %bb.0:
	push    rbp
	mov    rbp, rsp
	and    rsp, -16
	sub    rsp, 320
	vmovdqa    xmmword ptr [rsp + 48], xmm0
	vmovdqa    xmmword ptr [rsp + 32], xmm1
	vmovdqa    xmm1, xmmword ptr [rsp + 48]
	vmovdqa    xmm0, xmmword ptr [rsp + 32]
	vmovdqa    xmmword ptr [rsp + 112], xmm1
	vmovdqa    xmmword ptr [rsp + 96], xmm0
	vmovdqa    xmm0, xmmword ptr [rsp + 112]
	vmovdqa    xmm1, xmmword ptr [rsp + 96]
	vpmullw    xmm0, xmm0, xmm1
	vmovdqa    xmmword ptr [rsp + 16], xmm0
	vmovdqa    xmm0, xmmword ptr [rsp + 48]
	vmovdqa    xmmword ptr [rsp + 208], xmm0
	mov    dword ptr [rsp + 204], 8
	vmovdqa    xmm0, xmmword ptr [rsp + 208]
	mov    eax, dword ptr [rsp + 204]
	vmovd    xmm1, eax
	vpsrlw    xmm1, xmm0, xmm1
	vmovdqa    xmm0, xmmword ptr [rsp + 32]
	vmovdqa    xmmword ptr [rsp + 176], xmm0
	mov    dword ptr [rsp + 172], 8
	vmovdqa    xmm0, xmmword ptr [rsp + 176]
	mov    eax, dword ptr [rsp + 172]
	vmovd    xmm2, eax
	vpsrlw    xmm0, xmm0, xmm2
	vmovdqa    xmmword ptr [rsp + 80], xmm1
	vmovdqa    xmmword ptr [rsp + 64], xmm0
	vmovdqa    xmm0, xmmword ptr [rsp + 80]
	vmovdqa    xmm1, xmmword ptr [rsp + 64]
	vpmullw    xmm0, xmm0, xmm1
	vmovdqa    xmmword ptr [rsp], xmm0
	vmovdqa    xmm0, xmmword ptr [rsp]
	vmovdqa    xmmword ptr [rsp + 304], xmm0
	mov    dword ptr [rsp + 300], 8
	vmovdqa    xmm0, xmmword ptr [rsp + 304]
	mov    eax, dword ptr [rsp + 300]
	vmovd    xmm1, eax
	vpsllw    xmm1, xmm0, xmm1
	vmovdqa    xmm0, xmmword ptr [rsp + 16]
	vmovdqa    xmmword ptr [rsp + 272], xmm0
	mov    dword ptr [rsp + 268], 8
	vmovdqa    xmm0, xmmword ptr [rsp + 272]
	mov    eax, dword ptr [rsp + 268]
	vmovd    xmm2, eax
	vpsllw    xmm0, xmm0, xmm2
	vmovdqa    xmmword ptr [rsp + 144], xmm0
	mov    dword ptr [rsp + 140], 8
	vmovdqa    xmm0, xmmword ptr [rsp + 144]
	mov    eax, dword ptr [rsp + 140]
	vmovd    xmm2, eax
	vpsrlw    xmm0, xmm0, xmm2
	vmovdqa    xmmword ptr [rsp + 240], xmm1
	vmovdqa    xmmword ptr [rsp + 224], xmm0
	vmovdqa    xmm0, xmmword ptr [rsp + 240]
	vpor    xmm0, xmm0, xmmword ptr [rsp + 224]
	mov    rsp, rbp
	pop    rbp
	ret
.Lfunc_end0:
	.size    MultiplyBytes, .Lfunc_end0-MultiplyBytes
                                        # -- End function
	.ident    "clang version 13.0.1"
	.section    ".note.GNU-stack","",@progbits
	.addrsig
