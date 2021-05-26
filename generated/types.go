// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#cgo pkg-config: ${SRCDIR}/../filcrypto.pc
#include "../filcrypto.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

<<<<<<< HEAD
<<<<<<< HEAD
// FilBLSSignature as declared in filecoin-ffi/filcrypto.h:61
=======
// FilBLSSignature as declared in filecoin-ffi/filcrypto.h:59
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilBLSSignature as declared in filecoin-ffi/filcrypto.h:59
>>>>>>> d1f84bb (run the make cgo-gen)
type FilBLSSignature struct {
	Inner          [96]byte
	refa2ac09ba    *C.fil_BLSSignature
	allocsa2ac09ba interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilAggregateResponse as declared in filecoin-ffi/filcrypto.h:68
=======
// FilAggregateResponse as declared in filecoin-ffi/filcrypto.h:66
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilAggregateResponse as declared in filecoin-ffi/filcrypto.h:66
>>>>>>> d1f84bb (run the make cgo-gen)
type FilAggregateResponse struct {
	Signature      FilBLSSignature
	refb3efa36d    *C.fil_AggregateResponse
	allocsb3efa36d interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilClearCacheResponse as declared in filecoin-ffi/filcrypto.h:73
=======
// FilClearCacheResponse as declared in filecoin-ffi/filcrypto.h:71
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilClearCacheResponse as declared in filecoin-ffi/filcrypto.h:71
>>>>>>> d1f84bb (run the make cgo-gen)
type FilClearCacheResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	refa9a80400    *C.fil_ClearCacheResponse
	allocsa9a80400 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilZeroSignatureResponse as declared in filecoin-ffi/filcrypto.h:80
type FilZeroSignatureResponse struct {
	Signature      FilBLSSignature
	ref835a0405    *C.fil_ZeroSignatureResponse
	allocs835a0405 interface{}
}

// FilFauxRepResponse as declared in filecoin-ffi/filcrypto.h:86
=======
// FilFauxRepResponse as declared in filecoin-ffi/filcrypto.h:77
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilFauxRepResponse as declared in filecoin-ffi/filcrypto.h:77
>>>>>>> d1f84bb (run the make cgo-gen)
type FilFauxRepResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	Commitment     [32]byte
	refaa003f71    *C.fil_FauxRepResponse
	allocsaa003f71 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilFinalizeTicketResponse as declared in filecoin-ffi/filcrypto.h:92
=======
// FilFinalizeTicketResponse as declared in filecoin-ffi/filcrypto.h:83
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilFinalizeTicketResponse as declared in filecoin-ffi/filcrypto.h:83
>>>>>>> d1f84bb (run the make cgo-gen)
type FilFinalizeTicketResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	Ticket         [32]byte
	refb370fa86    *C.fil_FinalizeTicketResponse
	allocsb370fa86 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilGenerateDataCommitmentResponse as declared in filecoin-ffi/filcrypto.h:98
=======
// FilGenerateDataCommitmentResponse as declared in filecoin-ffi/filcrypto.h:89
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilGenerateDataCommitmentResponse as declared in filecoin-ffi/filcrypto.h:89
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGenerateDataCommitmentResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	CommD          [32]byte
	ref87da7dd9    *C.fil_GenerateDataCommitmentResponse
	allocs87da7dd9 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilGenerateFallbackSectorChallengesResponse as declared in filecoin-ffi/filcrypto.h:108
=======
// FilGenerateFallbackSectorChallengesResponse as declared in filecoin-ffi/filcrypto.h:99
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGenerateFallbackSectorChallengesResponse struct {
	ErrorMsg         string
	StatusCode       FCPResponseStatus
	IdsPtr           []uint64
	IdsLen           uint
	ChallengesPtr    []uint64
	ChallengesLen    uint
	ChallengesStride uint
	ref7047a3fa      *C.fil_GenerateFallbackSectorChallengesResponse
	allocs7047a3fa   interface{}
}

<<<<<<< HEAD
// FilGeneratePieceCommitmentResponse as declared in filecoin-ffi/filcrypto.h:119
=======
// FilGeneratePieceCommitmentResponse as declared in filecoin-ffi/filcrypto.h:100
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilGeneratePieceCommitmentResponse as declared in filecoin-ffi/filcrypto.h:110
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGeneratePieceCommitmentResponse struct {
	StatusCode      FCPResponseStatus
	ErrorMsg        string
	CommP           [32]byte
	NumBytesAligned uint64
	ref4b00fda4     *C.fil_GeneratePieceCommitmentResponse
	allocs4b00fda4  interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilVanillaProof as declared in filecoin-ffi/filcrypto.h:124
=======
// FilVanillaProof as declared in filecoin-ffi/filcrypto.h:115
>>>>>>> d1f84bb (run the make cgo-gen)
type FilVanillaProof struct {
	ProofLen       uint
	ProofPtr       []byte
	refb3e7638c    *C.fil_VanillaProof
	allocsb3e7638c interface{}
}

<<<<<<< HEAD
// FilGenerateSingleVanillaProofResponse as declared in filecoin-ffi/filcrypto.h:130
=======
// FilGenerateSingleVanillaProofResponse as declared in filecoin-ffi/filcrypto.h:121
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGenerateSingleVanillaProofResponse struct {
	ErrorMsg       string
	VanillaProof   FilVanillaProof
	StatusCode     FCPResponseStatus
	reff9d21b04    *C.fil_GenerateSingleVanillaProofResponse
	allocsf9d21b04 interface{}
}

<<<<<<< HEAD
// FilPoStProof as declared in filecoin-ffi/filcrypto.h:136
=======
// FilPoStProof as declared in filecoin-ffi/filcrypto.h:106
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPoStProof as declared in filecoin-ffi/filcrypto.h:127
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPoStProof struct {
	RegisteredProof FilRegisteredPoStProof
	ProofLen        uint
	ProofPtr        []byte
	ref3451bfa      *C.fil_PoStProof
	allocs3451bfa   interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilGenerateWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:145
=======
// FilGenerateWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:115
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilGenerateWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:136
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGenerateWindowPoStResponse struct {
	ErrorMsg         string
	ProofsLen        uint
	ProofsPtr        []FilPoStProof
	FaultySectorsLen uint
	FaultySectorsPtr []uint64
	StatusCode       FCPResponseStatus
	ref2a5f3ba8      *C.fil_GenerateWindowPoStResponse
	allocs2a5f3ba8   interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilGenerateWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:152
=======
// FilGenerateWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:122
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilGenerateWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:143
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGenerateWinningPoStResponse struct {
	ErrorMsg       string
	ProofsLen      uint
	ProofsPtr      []FilPoStProof
	StatusCode     FCPResponseStatus
	ref1405b8ec    *C.fil_GenerateWinningPoStResponse
	allocs1405b8ec interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilGenerateWinningPoStSectorChallenge as declared in filecoin-ffi/filcrypto.h:159
=======
// FilGenerateWinningPoStSectorChallenge as declared in filecoin-ffi/filcrypto.h:129
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilGenerateWinningPoStSectorChallenge as declared in filecoin-ffi/filcrypto.h:150
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGenerateWinningPoStSectorChallenge struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	IdsPtr         []uint64
	IdsLen         uint
	ref69d2a405    *C.fil_GenerateWinningPoStSectorChallenge
	allocs69d2a405 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilGpuDeviceResponse as declared in filecoin-ffi/filcrypto.h:166
=======
// FilGpuDeviceResponse as declared in filecoin-ffi/filcrypto.h:136
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilGpuDeviceResponse as declared in filecoin-ffi/filcrypto.h:157
>>>>>>> d1f84bb (run the make cgo-gen)
type FilGpuDeviceResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	DevicesLen     uint
	DevicesPtr     []string
	ref58f92915    *C.fil_GpuDeviceResponse
	allocs58f92915 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilBLSDigest as declared in filecoin-ffi/filcrypto.h:170
=======
// FilBLSDigest as declared in filecoin-ffi/filcrypto.h:140
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilBLSDigest as declared in filecoin-ffi/filcrypto.h:161
>>>>>>> d1f84bb (run the make cgo-gen)
type FilBLSDigest struct {
	Inner          [96]byte
	ref215fc78c    *C.fil_BLSDigest
	allocs215fc78c interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilHashResponse as declared in filecoin-ffi/filcrypto.h:177
=======
// FilHashResponse as declared in filecoin-ffi/filcrypto.h:147
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilHashResponse as declared in filecoin-ffi/filcrypto.h:168
>>>>>>> d1f84bb (run the make cgo-gen)
type FilHashResponse struct {
	Digest         FilBLSDigest
	refc52a22ef    *C.fil_HashResponse
	allocsc52a22ef interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilInitLogFdResponse as declared in filecoin-ffi/filcrypto.h:182
=======
// FilInitLogFdResponse as declared in filecoin-ffi/filcrypto.h:152
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilInitLogFdResponse as declared in filecoin-ffi/filcrypto.h:173
>>>>>>> d1f84bb (run the make cgo-gen)
type FilInitLogFdResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref3c1a0a08    *C.fil_InitLogFdResponse
	allocs3c1a0a08 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilBLSPrivateKey as declared in filecoin-ffi/filcrypto.h:186
=======
// FilBLSPrivateKey as declared in filecoin-ffi/filcrypto.h:156
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilBLSPrivateKey as declared in filecoin-ffi/filcrypto.h:177
>>>>>>> d1f84bb (run the make cgo-gen)
type FilBLSPrivateKey struct {
	Inner          [32]byte
	ref2f77fe3a    *C.fil_BLSPrivateKey
	allocs2f77fe3a interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilPrivateKeyGenerateResponse as declared in filecoin-ffi/filcrypto.h:193
=======
// FilPrivateKeyGenerateResponse as declared in filecoin-ffi/filcrypto.h:163
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPrivateKeyGenerateResponse as declared in filecoin-ffi/filcrypto.h:184
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPrivateKeyGenerateResponse struct {
	PrivateKey    FilBLSPrivateKey
	ref2dba09f    *C.fil_PrivateKeyGenerateResponse
	allocs2dba09f interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilBLSPublicKey as declared in filecoin-ffi/filcrypto.h:197
=======
// FilBLSPublicKey as declared in filecoin-ffi/filcrypto.h:167
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilBLSPublicKey as declared in filecoin-ffi/filcrypto.h:188
>>>>>>> d1f84bb (run the make cgo-gen)
type FilBLSPublicKey struct {
	Inner          [48]byte
	ref6d0cab13    *C.fil_BLSPublicKey
	allocs6d0cab13 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilPrivateKeyPublicKeyResponse as declared in filecoin-ffi/filcrypto.h:204
=======
// FilPrivateKeyPublicKeyResponse as declared in filecoin-ffi/filcrypto.h:174
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPrivateKeyPublicKeyResponse as declared in filecoin-ffi/filcrypto.h:195
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPrivateKeyPublicKeyResponse struct {
	PublicKey      FilBLSPublicKey
	refee14e59d    *C.fil_PrivateKeyPublicKeyResponse
	allocsee14e59d interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilPrivateKeySignResponse as declared in filecoin-ffi/filcrypto.h:211
=======
// FilPrivateKeySignResponse as declared in filecoin-ffi/filcrypto.h:181
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPrivateKeySignResponse as declared in filecoin-ffi/filcrypto.h:202
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPrivateKeySignResponse struct {
	Signature      FilBLSSignature
	refcdf97b28    *C.fil_PrivateKeySignResponse
	allocscdf97b28 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilSealCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:218
=======
// FilSealCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:188
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilSealCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:209
>>>>>>> d1f84bb (run the make cgo-gen)
type FilSealCommitPhase1Response struct {
	StatusCode                FCPResponseStatus
	ErrorMsg                  string
	SealCommitPhase1OutputPtr []byte
	SealCommitPhase1OutputLen uint
	ref61ed8561               *C.fil_SealCommitPhase1Response
	allocs61ed8561            interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilSealCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:225
=======
// FilSealCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:195
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilSealCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:216
>>>>>>> d1f84bb (run the make cgo-gen)
type FilSealCommitPhase2Response struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ProofPtr       []byte
	ProofLen       uint
	ref5860b9a4    *C.fil_SealCommitPhase2Response
	allocs5860b9a4 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilSealPreCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:232
=======
// FilSealPreCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:202
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilSealPreCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:223
>>>>>>> d1f84bb (run the make cgo-gen)
type FilSealPreCommitPhase1Response struct {
	ErrorMsg                     string
	StatusCode                   FCPResponseStatus
	SealPreCommitPhase1OutputPtr []byte
	SealPreCommitPhase1OutputLen uint
	ref132bbfd8                  *C.fil_SealPreCommitPhase1Response
	allocs132bbfd8               interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilSealPreCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:240
=======
// FilSealPreCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:210
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilSealPreCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:231
>>>>>>> d1f84bb (run the make cgo-gen)
type FilSealPreCommitPhase2Response struct {
	ErrorMsg        string
	StatusCode      FCPResponseStatus
	RegisteredProof FilRegisteredSealProof
	CommD           [32]byte
	CommR           [32]byte
	ref2aa6831d     *C.fil_SealPreCommitPhase2Response
	allocs2aa6831d  interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilStringResponse as declared in filecoin-ffi/filcrypto.h:249
=======
// FilStringResponse as declared in filecoin-ffi/filcrypto.h:219
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilStringResponse as declared in filecoin-ffi/filcrypto.h:240
>>>>>>> d1f84bb (run the make cgo-gen)
type FilStringResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	StringVal      string
	ref4f413043    *C.fil_StringResponse
	allocs4f413043 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilUnsealRangeResponse as declared in filecoin-ffi/filcrypto.h:254
=======
// FilUnsealRangeResponse as declared in filecoin-ffi/filcrypto.h:224
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilUnsealRangeResponse as declared in filecoin-ffi/filcrypto.h:245
>>>>>>> d1f84bb (run the make cgo-gen)
type FilUnsealRangeResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref61e219c9    *C.fil_UnsealRangeResponse
	allocs61e219c9 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilVerifySealResponse as declared in filecoin-ffi/filcrypto.h:260
=======
// FilVerifySealResponse as declared in filecoin-ffi/filcrypto.h:230
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilVerifySealResponse as declared in filecoin-ffi/filcrypto.h:251
>>>>>>> d1f84bb (run the make cgo-gen)
type FilVerifySealResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refd4397079    *C.fil_VerifySealResponse
	allocsd4397079 interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilVerifyWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:266
=======
// FilVerifyWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:236
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilVerifyWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:257
>>>>>>> d1f84bb (run the make cgo-gen)
type FilVerifyWindowPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	ref34c4d49f    *C.fil_VerifyWindowPoStResponse
	allocs34c4d49f interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilVerifyWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:272
=======
// FilVerifyWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:242
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilVerifyWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:263
>>>>>>> d1f84bb (run the make cgo-gen)
type FilVerifyWinningPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refaca6860c    *C.fil_VerifyWinningPoStResponse
	allocsaca6860c interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilWriteWithAlignmentResponse as declared in filecoin-ffi/filcrypto.h:280
=======
// FilWriteWithAlignmentResponse as declared in filecoin-ffi/filcrypto.h:250
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilWriteWithAlignmentResponse as declared in filecoin-ffi/filcrypto.h:271
>>>>>>> d1f84bb (run the make cgo-gen)
type FilWriteWithAlignmentResponse struct {
	CommP                 [32]byte
	ErrorMsg              string
	LeftAlignmentUnpadded uint64
	StatusCode            FCPResponseStatus
	TotalWriteUnpadded    uint64
	refa330e79            *C.fil_WriteWithAlignmentResponse
	allocsa330e79         interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilWriteWithoutAlignmentResponse as declared in filecoin-ffi/filcrypto.h:287
=======
// FilWriteWithoutAlignmentResponse as declared in filecoin-ffi/filcrypto.h:257
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilWriteWithoutAlignmentResponse as declared in filecoin-ffi/filcrypto.h:278
>>>>>>> d1f84bb (run the make cgo-gen)
type FilWriteWithoutAlignmentResponse struct {
	CommP              [32]byte
	ErrorMsg           string
	StatusCode         FCPResponseStatus
	TotalWriteUnpadded uint64
	refc8e1ed8         *C.fil_WriteWithoutAlignmentResponse
	allocsc8e1ed8      interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilPublicPieceInfo as declared in filecoin-ffi/filcrypto.h:292
=======
// FilPublicPieceInfo as declared in filecoin-ffi/filcrypto.h:262
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPublicPieceInfo as declared in filecoin-ffi/filcrypto.h:283
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPublicPieceInfo struct {
	NumBytes       uint64
	CommP          [32]byte
	refd00025ac    *C.fil_PublicPieceInfo
	allocsd00025ac interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// Fil32ByteArray as declared in filecoin-ffi/filcrypto.h:296
=======
// Fil32ByteArray as declared in filecoin-ffi/filcrypto.h:266
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// Fil32ByteArray as declared in filecoin-ffi/filcrypto.h:287
>>>>>>> d1f84bb (run the make cgo-gen)
type Fil32ByteArray struct {
	Inner          [32]byte
	ref373ec61a    *C.fil_32ByteArray
	allocs373ec61a interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilPrivateReplicaInfo as declared in filecoin-ffi/filcrypto.h:304
=======
// FilPrivateReplicaInfo as declared in filecoin-ffi/filcrypto.h:274
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPrivateReplicaInfo as declared in filecoin-ffi/filcrypto.h:295
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPrivateReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CacheDirPath    string
	CommR           [32]byte
	ReplicaPath     string
	SectorId        uint64
	ref81a31e9b     *C.fil_PrivateReplicaInfo
	allocs81a31e9b  interface{}
}

<<<<<<< HEAD
<<<<<<< HEAD
// FilPublicReplicaInfo as declared in filecoin-ffi/filcrypto.h:310
=======
// FilPublicReplicaInfo as declared in filecoin-ffi/filcrypto.h:280
>>>>>>> 803fdd6 (move 8GiB enum to the end)
=======
// FilPublicReplicaInfo as declared in filecoin-ffi/filcrypto.h:301
>>>>>>> d1f84bb (run the make cgo-gen)
type FilPublicReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CommR           [32]byte
	SectorId        uint64
	ref81b617c2     *C.fil_PublicReplicaInfo
	allocs81b617c2  interface{}
}
