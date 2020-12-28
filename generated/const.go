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

// FCPResponseStatus as declared in filecoin-ffi/filcrypto.h:31
type FCPResponseStatus int32

// FCPResponseStatus enumeration from filecoin-ffi/filcrypto.h:31
const (
	FCPResponseStatusFCPNoError           FCPResponseStatus = C.FCPResponseStatus_FCPNoError
	FCPResponseStatusFCPUnclassifiedError FCPResponseStatus = C.FCPResponseStatus_FCPUnclassifiedError
	FCPResponseStatusFCPCallerError       FCPResponseStatus = C.FCPResponseStatus_FCPCallerError
	FCPResponseStatusFCPReceiverError     FCPResponseStatus = C.FCPResponseStatus_FCPReceiverError
)

// FilRegisteredPoStProof as declared in filecoin-ffi/filcrypto.h:46
type FilRegisteredPoStProof int32

// FilRegisteredPoStProof enumeration from filecoin-ffi/filcrypto.h:46
const (
	FilRegisteredPoStProofStackedDrgWinning2KiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning2KiBV1
	FilRegisteredPoStProofStackedDrgWinning8MiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning8MiBV1
	FilRegisteredPoStProofStackedDrgWinning512MiBV1 FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning512MiBV1
	FilRegisteredPoStProofStackedDrgWinning32GiBV1  FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning32GiBV1
	FilRegisteredPoStProofStackedDrgWinning64GiBV1  FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning64GiBV1
	FilRegisteredPoStProofStackedDrgWindow2KiBV1    FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow2KiBV1
	FilRegisteredPoStProofStackedDrgWindow8MiBV1    FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow8MiBV1
	FilRegisteredPoStProofStackedDrgWindow512MiBV1  FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow512MiBV1
	FilRegisteredPoStProofStackedDrgWindow32GiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow32GiBV1
	FilRegisteredPoStProofStackedDrgWindow64GiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow64GiBV1
	FilRegisteredPoStProofStackedDrgWinning8GiBV1   FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWinning8GiBV1
	FilRegisteredPoStProofStackedDrgWindow8GiBV1    FilRegisteredPoStProof = C.fil_RegisteredPoStProof_StackedDrgWindow8GiBV1
)

<<<<<<< HEAD
// FilRegisteredSealProof as declared in filecoin-ffi/filcrypto.h:57
type FilRegisteredSealProof int32

// FilRegisteredSealProof enumeration from filecoin-ffi/filcrypto.h:57
=======
// FilRegisteredSealProof as declared in filecoin-ffi/filcrypto.h:55
type FilRegisteredSealProof int32

// FilRegisteredSealProof enumeration from filecoin-ffi/filcrypto.h:55
>>>>>>> 803fdd6 (move 8GiB enum to the end)
const (
<<<<<<< HEAD
	FilRegisteredSealProofStackedDrg2KiBV1    FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg2KiBV1
	FilRegisteredSealProofStackedDrg8MiBV1    FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg8MiBV1
	FilRegisteredSealProofStackedDrg512MiBV1  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg512MiBV1
	FilRegisteredSealProofStackedDrg32GiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg32GiBV1
	FilRegisteredSealProofStackedDrg64GiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg64GiBV1
	FilRegisteredSealProofStackedDrg2KiBV11   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg2KiBV1_1
	FilRegisteredSealProofStackedDrg8MiBV11   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg8MiBV1_1
	FilRegisteredSealProofStackedDrg512MiBV11 FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg512MiBV1_1
	FilRegisteredSealProofStackedDrg32GiBV11  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg32GiBV1_1
	FilRegisteredSealProofStackedDrg64GiBV11  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg64GiBV1_1
=======
	FilRegisteredSealProofStackedDrg2KiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg2KiBV1
	FilRegisteredSealProofStackedDrg8MiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg8MiBV1
	FilRegisteredSealProofStackedDrg512MiBV1 FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg512MiBV1
	FilRegisteredSealProofStackedDrg32GiBV1  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg32GiBV1
	FilRegisteredSealProofStackedDrg64GiBV1  FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg64GiBV1
<<<<<<< HEAD
>>>>>>> fdaca8e (update Cargo.toml and add 8G 5 lays)
=======
	FilRegisteredSealProofStackedDrg8GiBV1   FilRegisteredSealProof = C.fil_RegisteredSealProof_StackedDrg8GiBV1
>>>>>>> 803fdd6 (move 8GiB enum to the end)
)
