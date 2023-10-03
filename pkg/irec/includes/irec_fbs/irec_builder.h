#ifndef IREC_BUILDER_H
#define IREC_BUILDER_H

/* Generated by flatcc 0.6.2 FlatBuffers schema compiler for C by dvide.com */

#ifndef IREC_READER_H
#include "irec_reader.h"
#endif
#ifndef FLATBUFFERS_COMMON_BUILDER_H
#include "flatbuffers_common_builder.h"
#endif
#include "flatcc/flatcc_prologue.h"
#ifndef flatbuffers_identifier
#define flatbuffers_identifier 0
#endif
#ifndef flatbuffers_extension
#define flatbuffers_extension "bin"
#endif

#define __IREC_Interface_formal_args , uint32_t v0
#define __IREC_Interface_call_args , v0
static inline IREC_Interface_t *IREC_Interface_assign(IREC_Interface_t *p, uint32_t v0)
{ p->id = v0;
  return p; }
static inline IREC_Interface_t *IREC_Interface_copy(IREC_Interface_t *p, const IREC_Interface_t *p2)
{ p->id = p2->id;
  return p; }
static inline IREC_Interface_t *IREC_Interface_assign_to_pe(IREC_Interface_t *p, uint32_t v0)
{ flatbuffers_uint32_assign_to_pe(&p->id, v0);
  return p; }
static inline IREC_Interface_t *IREC_Interface_copy_to_pe(IREC_Interface_t *p, const IREC_Interface_t *p2)
{ flatbuffers_uint32_copy_to_pe(&p->id, &p2->id);
  return p; }
static inline IREC_Interface_t *IREC_Interface_assign_from_pe(IREC_Interface_t *p, uint32_t v0)
{ flatbuffers_uint32_assign_from_pe(&p->id, v0);
  return p; }
static inline IREC_Interface_t *IREC_Interface_copy_from_pe(IREC_Interface_t *p, const IREC_Interface_t *p2)
{ flatbuffers_uint32_copy_from_pe(&p->id, &p2->id);
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_Interface, 4, 4, IREC_Interface_file_identifier, IREC_Interface_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_Interface, IREC_Interface_t)

#define __IREC_MapEntryUlongUint_formal_args , uint64_t v0, uint32_t v1
#define __IREC_MapEntryUlongUint_call_args , v0, v1
static inline IREC_MapEntryUlongUint_t *IREC_MapEntryUlongUint_assign(IREC_MapEntryUlongUint_t *p, uint64_t v0, uint32_t v1)
{ p->k = v0; p->v = v1;
  return p; }
static inline IREC_MapEntryUlongUint_t *IREC_MapEntryUlongUint_copy(IREC_MapEntryUlongUint_t *p, const IREC_MapEntryUlongUint_t *p2)
{ p->k = p2->k; p->v = p2->v;
  return p; }
static inline IREC_MapEntryUlongUint_t *IREC_MapEntryUlongUint_assign_to_pe(IREC_MapEntryUlongUint_t *p, uint64_t v0, uint32_t v1)
{ flatbuffers_uint64_assign_to_pe(&p->k, v0); flatbuffers_uint32_assign_to_pe(&p->v, v1);
  return p; }
static inline IREC_MapEntryUlongUint_t *IREC_MapEntryUlongUint_copy_to_pe(IREC_MapEntryUlongUint_t *p, const IREC_MapEntryUlongUint_t *p2)
{ flatbuffers_uint64_copy_to_pe(&p->k, &p2->k); flatbuffers_uint32_copy_to_pe(&p->v, &p2->v);
  return p; }
static inline IREC_MapEntryUlongUint_t *IREC_MapEntryUlongUint_assign_from_pe(IREC_MapEntryUlongUint_t *p, uint64_t v0, uint32_t v1)
{ flatbuffers_uint64_assign_from_pe(&p->k, v0); flatbuffers_uint32_assign_from_pe(&p->v, v1);
  return p; }
static inline IREC_MapEntryUlongUint_t *IREC_MapEntryUlongUint_copy_from_pe(IREC_MapEntryUlongUint_t *p, const IREC_MapEntryUlongUint_t *p2)
{ flatbuffers_uint64_copy_from_pe(&p->k, &p2->k); flatbuffers_uint32_copy_from_pe(&p->v, &p2->v);
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_MapEntryUlongUint, 16, 8, IREC_MapEntryUlongUint_file_identifier, IREC_MapEntryUlongUint_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_MapEntryUlongUint, IREC_MapEntryUlongUint_t)

#define __IREC_MapEntryUlongUlong_formal_args , uint64_t v0, uint64_t v1
#define __IREC_MapEntryUlongUlong_call_args , v0, v1
static inline IREC_MapEntryUlongUlong_t *IREC_MapEntryUlongUlong_assign(IREC_MapEntryUlongUlong_t *p, uint64_t v0, uint64_t v1)
{ p->k = v0; p->v = v1;
  return p; }
static inline IREC_MapEntryUlongUlong_t *IREC_MapEntryUlongUlong_copy(IREC_MapEntryUlongUlong_t *p, const IREC_MapEntryUlongUlong_t *p2)
{ p->k = p2->k; p->v = p2->v;
  return p; }
static inline IREC_MapEntryUlongUlong_t *IREC_MapEntryUlongUlong_assign_to_pe(IREC_MapEntryUlongUlong_t *p, uint64_t v0, uint64_t v1)
{ flatbuffers_uint64_assign_to_pe(&p->k, v0); flatbuffers_uint64_assign_to_pe(&p->v, v1);
  return p; }
static inline IREC_MapEntryUlongUlong_t *IREC_MapEntryUlongUlong_copy_to_pe(IREC_MapEntryUlongUlong_t *p, const IREC_MapEntryUlongUlong_t *p2)
{ flatbuffers_uint64_copy_to_pe(&p->k, &p2->k); flatbuffers_uint64_copy_to_pe(&p->v, &p2->v);
  return p; }
static inline IREC_MapEntryUlongUlong_t *IREC_MapEntryUlongUlong_assign_from_pe(IREC_MapEntryUlongUlong_t *p, uint64_t v0, uint64_t v1)
{ flatbuffers_uint64_assign_from_pe(&p->k, v0); flatbuffers_uint64_assign_from_pe(&p->v, v1);
  return p; }
static inline IREC_MapEntryUlongUlong_t *IREC_MapEntryUlongUlong_copy_from_pe(IREC_MapEntryUlongUlong_t *p, const IREC_MapEntryUlongUlong_t *p2)
{ flatbuffers_uint64_copy_from_pe(&p->k, &p2->k); flatbuffers_uint64_copy_from_pe(&p->v, &p2->v);
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_MapEntryUlongUlong, 16, 8, IREC_MapEntryUlongUlong_file_identifier, IREC_MapEntryUlongUlong_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_MapEntryUlongUlong, IREC_MapEntryUlongUlong_t)

#define __IREC_HiddenPathExt_formal_args , flatbuffers_bool_t v0
#define __IREC_HiddenPathExt_call_args , v0
static inline IREC_HiddenPathExt_t *IREC_HiddenPathExt_assign(IREC_HiddenPathExt_t *p, flatbuffers_bool_t v0)
{ p->is_hidden = v0;
  return p; }
static inline IREC_HiddenPathExt_t *IREC_HiddenPathExt_copy(IREC_HiddenPathExt_t *p, const IREC_HiddenPathExt_t *p2)
{ p->is_hidden = p2->is_hidden;
  return p; }
static inline IREC_HiddenPathExt_t *IREC_HiddenPathExt_assign_to_pe(IREC_HiddenPathExt_t *p, flatbuffers_bool_t v0)
{ p->is_hidden = v0;
  return p; }
static inline IREC_HiddenPathExt_t *IREC_HiddenPathExt_copy_to_pe(IREC_HiddenPathExt_t *p, const IREC_HiddenPathExt_t *p2)
{ p->is_hidden = p2->is_hidden;
  return p; }
static inline IREC_HiddenPathExt_t *IREC_HiddenPathExt_assign_from_pe(IREC_HiddenPathExt_t *p, flatbuffers_bool_t v0)
{ p->is_hidden = v0;
  return p; }
static inline IREC_HiddenPathExt_t *IREC_HiddenPathExt_copy_from_pe(IREC_HiddenPathExt_t *p, const IREC_HiddenPathExt_t *p2)
{ p->is_hidden = p2->is_hidden;
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_HiddenPathExt, 1, 1, IREC_HiddenPathExt_file_identifier, IREC_HiddenPathExt_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_HiddenPathExt, IREC_HiddenPathExt_t)

#define __IREC_IrecExt_formal_args , uint32_t v0
#define __IREC_IrecExt_call_args , v0
static inline IREC_IrecExt_t *IREC_IrecExt_assign(IREC_IrecExt_t *p, uint32_t v0)
{ p->intf_group = v0;
  return p; }
static inline IREC_IrecExt_t *IREC_IrecExt_copy(IREC_IrecExt_t *p, const IREC_IrecExt_t *p2)
{ p->intf_group = p2->intf_group;
  return p; }
static inline IREC_IrecExt_t *IREC_IrecExt_assign_to_pe(IREC_IrecExt_t *p, uint32_t v0)
{ flatbuffers_uint32_assign_to_pe(&p->intf_group, v0);
  return p; }
static inline IREC_IrecExt_t *IREC_IrecExt_copy_to_pe(IREC_IrecExt_t *p, const IREC_IrecExt_t *p2)
{ flatbuffers_uint32_copy_to_pe(&p->intf_group, &p2->intf_group);
  return p; }
static inline IREC_IrecExt_t *IREC_IrecExt_assign_from_pe(IREC_IrecExt_t *p, uint32_t v0)
{ flatbuffers_uint32_assign_from_pe(&p->intf_group, v0);
  return p; }
static inline IREC_IrecExt_t *IREC_IrecExt_copy_from_pe(IREC_IrecExt_t *p, const IREC_IrecExt_t *p2)
{ flatbuffers_uint32_copy_from_pe(&p->intf_group, &p2->intf_group);
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_IrecExt, 4, 4, IREC_IrecExt_file_identifier, IREC_IrecExt_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_IrecExt, IREC_IrecExt_t)

#define __IREC_HopField_formal_args , uint64_t v0, uint64_t v1, uint32_t v2
#define __IREC_HopField_call_args , v0, v1, v2
static inline IREC_HopField_t *IREC_HopField_assign(IREC_HopField_t *p, uint64_t v0, uint64_t v1, uint32_t v2)
{ p->ingress = v0; p->egress = v1; p->exp_time = v2;
  return p; }
static inline IREC_HopField_t *IREC_HopField_copy(IREC_HopField_t *p, const IREC_HopField_t *p2)
{ p->ingress = p2->ingress; p->egress = p2->egress; p->exp_time = p2->exp_time;
  return p; }
static inline IREC_HopField_t *IREC_HopField_assign_to_pe(IREC_HopField_t *p, uint64_t v0, uint64_t v1, uint32_t v2)
{ flatbuffers_uint64_assign_to_pe(&p->ingress, v0); flatbuffers_uint64_assign_to_pe(&p->egress, v1); flatbuffers_uint32_assign_to_pe(&p->exp_time, v2);
  return p; }
static inline IREC_HopField_t *IREC_HopField_copy_to_pe(IREC_HopField_t *p, const IREC_HopField_t *p2)
{ flatbuffers_uint64_copy_to_pe(&p->ingress, &p2->ingress); flatbuffers_uint64_copy_to_pe(&p->egress, &p2->egress); flatbuffers_uint32_copy_to_pe(&p->exp_time, &p2->exp_time);
  return p; }
static inline IREC_HopField_t *IREC_HopField_assign_from_pe(IREC_HopField_t *p, uint64_t v0, uint64_t v1, uint32_t v2)
{ flatbuffers_uint64_assign_from_pe(&p->ingress, v0); flatbuffers_uint64_assign_from_pe(&p->egress, v1); flatbuffers_uint32_assign_from_pe(&p->exp_time, v2);
  return p; }
static inline IREC_HopField_t *IREC_HopField_copy_from_pe(IREC_HopField_t *p, const IREC_HopField_t *p2)
{ flatbuffers_uint64_copy_from_pe(&p->ingress, &p2->ingress); flatbuffers_uint64_copy_from_pe(&p->egress, &p2->egress); flatbuffers_uint32_copy_from_pe(&p->exp_time, &p2->exp_time);
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_HopField, 24, 8, IREC_HopField_file_identifier, IREC_HopField_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_HopField, IREC_HopField_t)

#define __IREC_PeerEntry_formal_args ,\
  uint64_t v0, uint64_t v1, uint32_t v2, uint64_t v3, uint64_t v4, uint32_t v5
#define __IREC_PeerEntry_call_args ,\
  v0, v1, v2, v3, v4, v5
static inline IREC_PeerEntry_t *IREC_PeerEntry_assign(IREC_PeerEntry_t *p,
  uint64_t v0, uint64_t v1, uint32_t v2, uint64_t v3, uint64_t v4, uint32_t v5)
{ p->peer_isd_as = v0; p->peer_interface = v1; p->peer_mtu = v2; IREC_HopField_assign(&p->hop_field,
  v3, v4, v5);
  return p; }
static inline IREC_PeerEntry_t *IREC_PeerEntry_copy(IREC_PeerEntry_t *p, const IREC_PeerEntry_t *p2)
{ p->peer_isd_as = p2->peer_isd_as; p->peer_interface = p2->peer_interface; p->peer_mtu = p2->peer_mtu; IREC_HopField_copy(&p->hop_field, &p2->hop_field);
  return p; }
static inline IREC_PeerEntry_t *IREC_PeerEntry_assign_to_pe(IREC_PeerEntry_t *p,
  uint64_t v0, uint64_t v1, uint32_t v2, uint64_t v3, uint64_t v4, uint32_t v5)
{ flatbuffers_uint64_assign_to_pe(&p->peer_isd_as, v0); flatbuffers_uint64_assign_to_pe(&p->peer_interface, v1); flatbuffers_uint32_assign_to_pe(&p->peer_mtu, v2); IREC_HopField_assign_to_pe(&p->hop_field,
  v3, v4, v5);
  return p; }
static inline IREC_PeerEntry_t *IREC_PeerEntry_copy_to_pe(IREC_PeerEntry_t *p, const IREC_PeerEntry_t *p2)
{ flatbuffers_uint64_copy_to_pe(&p->peer_isd_as, &p2->peer_isd_as); flatbuffers_uint64_copy_to_pe(&p->peer_interface, &p2->peer_interface); flatbuffers_uint32_copy_to_pe(&p->peer_mtu, &p2->peer_mtu); IREC_HopField_copy_to_pe(&p->hop_field, &p2->hop_field);
  return p; }
static inline IREC_PeerEntry_t *IREC_PeerEntry_assign_from_pe(IREC_PeerEntry_t *p,
  uint64_t v0, uint64_t v1, uint32_t v2, uint64_t v3, uint64_t v4, uint32_t v5)
{ flatbuffers_uint64_assign_from_pe(&p->peer_isd_as, v0); flatbuffers_uint64_assign_from_pe(&p->peer_interface, v1); flatbuffers_uint32_assign_from_pe(&p->peer_mtu, v2); IREC_HopField_assign_from_pe(&p->hop_field,
  v3, v4, v5);
  return p; }
static inline IREC_PeerEntry_t *IREC_PeerEntry_copy_from_pe(IREC_PeerEntry_t *p, const IREC_PeerEntry_t *p2)
{ flatbuffers_uint64_copy_from_pe(&p->peer_isd_as, &p2->peer_isd_as); flatbuffers_uint64_copy_from_pe(&p->peer_interface, &p2->peer_interface); flatbuffers_uint32_copy_from_pe(&p->peer_mtu, &p2->peer_mtu); IREC_HopField_copy_from_pe(&p->hop_field, &p2->hop_field);
  return p; }
__flatbuffers_build_struct(flatbuffers_, IREC_PeerEntry, 48, 8, IREC_PeerEntry_file_identifier, IREC_PeerEntry_type_identifier)
__flatbuffers_define_fixed_array_primitives(flatbuffers_, IREC_PeerEntry, IREC_PeerEntry_t)

static const flatbuffers_voffset_t __IREC_LatencyInfo_required[] = { 0 };
typedef flatbuffers_ref_t IREC_LatencyInfo_ref_t;
static IREC_LatencyInfo_ref_t IREC_LatencyInfo_clone(flatbuffers_builder_t *B, IREC_LatencyInfo_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_LatencyInfo, 2)

static const flatbuffers_voffset_t __IREC_BandwidthInfo_required[] = { 0 };
typedef flatbuffers_ref_t IREC_BandwidthInfo_ref_t;
static IREC_BandwidthInfo_ref_t IREC_BandwidthInfo_clone(flatbuffers_builder_t *B, IREC_BandwidthInfo_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_BandwidthInfo, 2)

static const flatbuffers_voffset_t __IREC_GeoCoordinates_required[] = { 0 };
typedef flatbuffers_ref_t IREC_GeoCoordinates_ref_t;
static IREC_GeoCoordinates_ref_t IREC_GeoCoordinates_clone(flatbuffers_builder_t *B, IREC_GeoCoordinates_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_GeoCoordinates, 4)

static const flatbuffers_voffset_t __IREC_StaticInfoExt_required[] = { 0 };
typedef flatbuffers_ref_t IREC_StaticInfoExt_ref_t;
static IREC_StaticInfoExt_ref_t IREC_StaticInfoExt_clone(flatbuffers_builder_t *B, IREC_StaticInfoExt_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_StaticInfoExt, 6)

static const flatbuffers_voffset_t __IREC_SignedExtensions_required[] = { 0 };
typedef flatbuffers_ref_t IREC_SignedExtensions_ref_t;
static IREC_SignedExtensions_ref_t IREC_SignedExtensions_clone(flatbuffers_builder_t *B, IREC_SignedExtensions_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_SignedExtensions, 3)

static const flatbuffers_voffset_t __IREC_UnsignedExtensions_required[] = { 0 };
typedef flatbuffers_ref_t IREC_UnsignedExtensions_ref_t;
static IREC_UnsignedExtensions_ref_t IREC_UnsignedExtensions_clone(flatbuffers_builder_t *B, IREC_UnsignedExtensions_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_UnsignedExtensions, 0)

static const flatbuffers_voffset_t __IREC_ASEntry_required[] = { 0 };
typedef flatbuffers_ref_t IREC_ASEntry_ref_t;
static IREC_ASEntry_ref_t IREC_ASEntry_clone(flatbuffers_builder_t *B, IREC_ASEntry_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_ASEntry, 8)

static const flatbuffers_voffset_t __IREC_Beacon_required[] = { 0 };
typedef flatbuffers_ref_t IREC_Beacon_ref_t;
static IREC_Beacon_ref_t IREC_Beacon_clone(flatbuffers_builder_t *B, IREC_Beacon_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_Beacon, 4)

static const flatbuffers_voffset_t __IREC_Execution_required[] = { 0 };
typedef flatbuffers_ref_t IREC_Execution_ref_t;
static IREC_Execution_ref_t IREC_Execution_clone(flatbuffers_builder_t *B, IREC_Execution_table_t t);
__flatbuffers_build_table(flatbuffers_, IREC_Execution, 3)

#define __IREC_LatencyInfo_formal_args , IREC_MapEntryUlongUint_vec_ref_t v0, IREC_MapEntryUlongUint_vec_ref_t v1
#define __IREC_LatencyInfo_call_args , v0, v1
static inline IREC_LatencyInfo_ref_t IREC_LatencyInfo_create(flatbuffers_builder_t *B __IREC_LatencyInfo_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_LatencyInfo, IREC_LatencyInfo_file_identifier, IREC_LatencyInfo_type_identifier)

#define __IREC_BandwidthInfo_formal_args , IREC_MapEntryUlongUlong_vec_ref_t v0, IREC_MapEntryUlongUlong_vec_ref_t v1
#define __IREC_BandwidthInfo_call_args , v0, v1
static inline IREC_BandwidthInfo_ref_t IREC_BandwidthInfo_create(flatbuffers_builder_t *B __IREC_BandwidthInfo_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_BandwidthInfo, IREC_BandwidthInfo_file_identifier, IREC_BandwidthInfo_type_identifier)

#define __IREC_GeoCoordinates_formal_args , uint64_t v0, float v1, float v2, flatbuffers_string_ref_t v3
#define __IREC_GeoCoordinates_call_args , v0, v1, v2, v3
static inline IREC_GeoCoordinates_ref_t IREC_GeoCoordinates_create(flatbuffers_builder_t *B __IREC_GeoCoordinates_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_GeoCoordinates, IREC_GeoCoordinates_file_identifier, IREC_GeoCoordinates_type_identifier)

#define __IREC_StaticInfoExt_formal_args ,\
  IREC_LatencyInfo_ref_t v0, IREC_BandwidthInfo_ref_t v1, IREC_GeoCoordinates_vec_ref_t v2, IREC_MapEntryUlongUint_vec_ref_t v3, IREC_MapEntryUlongUint_vec_ref_t v4, flatbuffers_string_ref_t v5
#define __IREC_StaticInfoExt_call_args ,\
  v0, v1, v2, v3, v4, v5
static inline IREC_StaticInfoExt_ref_t IREC_StaticInfoExt_create(flatbuffers_builder_t *B __IREC_StaticInfoExt_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_StaticInfoExt, IREC_StaticInfoExt_file_identifier, IREC_StaticInfoExt_type_identifier)

#define __IREC_SignedExtensions_formal_args , IREC_StaticInfoExt_ref_t v0, IREC_HiddenPathExt_t *v1, IREC_IrecExt_t *v2
#define __IREC_SignedExtensions_call_args , v0, v1, v2
static inline IREC_SignedExtensions_ref_t IREC_SignedExtensions_create(flatbuffers_builder_t *B __IREC_SignedExtensions_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_SignedExtensions, IREC_SignedExtensions_file_identifier, IREC_SignedExtensions_type_identifier)

#define __IREC_UnsignedExtensions_formal_args 
#define __IREC_UnsignedExtensions_call_args 
static inline IREC_UnsignedExtensions_ref_t IREC_UnsignedExtensions_create(flatbuffers_builder_t *B __IREC_UnsignedExtensions_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_UnsignedExtensions, IREC_UnsignedExtensions_file_identifier, IREC_UnsignedExtensions_type_identifier)

#define __IREC_ASEntry_formal_args ,\
  uint64_t v0, uint64_t v1, uint32_t v2, IREC_HopField_t *v3,\
  IREC_PeerEntry_vec_ref_t v4, uint32_t v5, IREC_SignedExtensions_ref_t v6, IREC_UnsignedExtensions_ref_t v7
#define __IREC_ASEntry_call_args ,\
  v0, v1, v2, v3,\
  v4, v5, v6, v7
static inline IREC_ASEntry_ref_t IREC_ASEntry_create(flatbuffers_builder_t *B __IREC_ASEntry_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_ASEntry, IREC_ASEntry_file_identifier, IREC_ASEntry_type_identifier)

#define __IREC_Beacon_formal_args , flatbuffers_int8_vec_ref_t v0, IREC_ASEntry_vec_ref_t v1, uint32_t v2, uint32_t v3
#define __IREC_Beacon_call_args , v0, v1, v2, v3
static inline IREC_Beacon_ref_t IREC_Beacon_create(flatbuffers_builder_t *B __IREC_Beacon_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_Beacon, IREC_Beacon_file_identifier, IREC_Beacon_type_identifier)

#define __IREC_Execution_formal_args , IREC_Beacon_vec_ref_t v0, flatbuffers_int8_vec_ref_t v1, IREC_Interface_vec_ref_t v2
#define __IREC_Execution_call_args , v0, v1, v2
static inline IREC_Execution_ref_t IREC_Execution_create(flatbuffers_builder_t *B __IREC_Execution_formal_args);
__flatbuffers_build_table_prolog(flatbuffers_, IREC_Execution, IREC_Execution_file_identifier, IREC_Execution_type_identifier)

__flatbuffers_build_vector_field(0, flatbuffers_, IREC_LatencyInfo_intra, IREC_MapEntryUlongUint, IREC_MapEntryUlongUint_t, IREC_LatencyInfo)
__flatbuffers_build_vector_field(1, flatbuffers_, IREC_LatencyInfo_inter, IREC_MapEntryUlongUint, IREC_MapEntryUlongUint_t, IREC_LatencyInfo)

static inline IREC_LatencyInfo_ref_t IREC_LatencyInfo_create(flatbuffers_builder_t *B __IREC_LatencyInfo_formal_args)
{
    if (IREC_LatencyInfo_start(B)
        || IREC_LatencyInfo_intra_add(B, v0)
        || IREC_LatencyInfo_inter_add(B, v1)) {
        return 0;
    }
    return IREC_LatencyInfo_end(B);
}

static IREC_LatencyInfo_ref_t IREC_LatencyInfo_clone(flatbuffers_builder_t *B, IREC_LatencyInfo_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_LatencyInfo_start(B)
        || IREC_LatencyInfo_intra_pick(B, t)
        || IREC_LatencyInfo_inter_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_LatencyInfo_end(B));
}

__flatbuffers_build_vector_field(0, flatbuffers_, IREC_BandwidthInfo_intra, IREC_MapEntryUlongUlong, IREC_MapEntryUlongUlong_t, IREC_BandwidthInfo)
__flatbuffers_build_vector_field(1, flatbuffers_, IREC_BandwidthInfo_inter, IREC_MapEntryUlongUlong, IREC_MapEntryUlongUlong_t, IREC_BandwidthInfo)

static inline IREC_BandwidthInfo_ref_t IREC_BandwidthInfo_create(flatbuffers_builder_t *B __IREC_BandwidthInfo_formal_args)
{
    if (IREC_BandwidthInfo_start(B)
        || IREC_BandwidthInfo_intra_add(B, v0)
        || IREC_BandwidthInfo_inter_add(B, v1)) {
        return 0;
    }
    return IREC_BandwidthInfo_end(B);
}

static IREC_BandwidthInfo_ref_t IREC_BandwidthInfo_clone(flatbuffers_builder_t *B, IREC_BandwidthInfo_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_BandwidthInfo_start(B)
        || IREC_BandwidthInfo_intra_pick(B, t)
        || IREC_BandwidthInfo_inter_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_BandwidthInfo_end(B));
}

__flatbuffers_build_scalar_field(0, flatbuffers_, IREC_GeoCoordinates_ifId, flatbuffers_uint64, uint64_t, 8, 8, UINT64_C(0), IREC_GeoCoordinates)
__flatbuffers_build_scalar_field(1, flatbuffers_, IREC_GeoCoordinates_latitude, flatbuffers_float, float, 4, 4, 0.00000000f, IREC_GeoCoordinates)
__flatbuffers_build_scalar_field(2, flatbuffers_, IREC_GeoCoordinates_longitude, flatbuffers_float, float, 4, 4, 0.00000000f, IREC_GeoCoordinates)
__flatbuffers_build_string_field(3, flatbuffers_, IREC_GeoCoordinates_address, IREC_GeoCoordinates)

static inline IREC_GeoCoordinates_ref_t IREC_GeoCoordinates_create(flatbuffers_builder_t *B __IREC_GeoCoordinates_formal_args)
{
    if (IREC_GeoCoordinates_start(B)
        || IREC_GeoCoordinates_ifId_add(B, v0)
        || IREC_GeoCoordinates_latitude_add(B, v1)
        || IREC_GeoCoordinates_longitude_add(B, v2)
        || IREC_GeoCoordinates_address_add(B, v3)) {
        return 0;
    }
    return IREC_GeoCoordinates_end(B);
}

static IREC_GeoCoordinates_ref_t IREC_GeoCoordinates_clone(flatbuffers_builder_t *B, IREC_GeoCoordinates_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_GeoCoordinates_start(B)
        || IREC_GeoCoordinates_ifId_pick(B, t)
        || IREC_GeoCoordinates_latitude_pick(B, t)
        || IREC_GeoCoordinates_longitude_pick(B, t)
        || IREC_GeoCoordinates_address_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_GeoCoordinates_end(B));
}

__flatbuffers_build_table_field(0, flatbuffers_, IREC_StaticInfoExt_latency, IREC_LatencyInfo, IREC_StaticInfoExt)
__flatbuffers_build_table_field(1, flatbuffers_, IREC_StaticInfoExt_bandwidth, IREC_BandwidthInfo, IREC_StaticInfoExt)
__flatbuffers_build_table_vector_field(2, flatbuffers_, IREC_StaticInfoExt_geo, IREC_GeoCoordinates, IREC_StaticInfoExt)
__flatbuffers_build_vector_field(3, flatbuffers_, IREC_StaticInfoExt_link_type, IREC_MapEntryUlongUint, IREC_MapEntryUlongUint_t, IREC_StaticInfoExt)
__flatbuffers_build_vector_field(4, flatbuffers_, IREC_StaticInfoExt_internal_hops, IREC_MapEntryUlongUint, IREC_MapEntryUlongUint_t, IREC_StaticInfoExt)
__flatbuffers_build_string_field(5, flatbuffers_, IREC_StaticInfoExt_note, IREC_StaticInfoExt)

static inline IREC_StaticInfoExt_ref_t IREC_StaticInfoExt_create(flatbuffers_builder_t *B __IREC_StaticInfoExt_formal_args)
{
    if (IREC_StaticInfoExt_start(B)
        || IREC_StaticInfoExt_latency_add(B, v0)
        || IREC_StaticInfoExt_bandwidth_add(B, v1)
        || IREC_StaticInfoExt_geo_add(B, v2)
        || IREC_StaticInfoExt_link_type_add(B, v3)
        || IREC_StaticInfoExt_internal_hops_add(B, v4)
        || IREC_StaticInfoExt_note_add(B, v5)) {
        return 0;
    }
    return IREC_StaticInfoExt_end(B);
}

static IREC_StaticInfoExt_ref_t IREC_StaticInfoExt_clone(flatbuffers_builder_t *B, IREC_StaticInfoExt_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_StaticInfoExt_start(B)
        || IREC_StaticInfoExt_latency_pick(B, t)
        || IREC_StaticInfoExt_bandwidth_pick(B, t)
        || IREC_StaticInfoExt_geo_pick(B, t)
        || IREC_StaticInfoExt_link_type_pick(B, t)
        || IREC_StaticInfoExt_internal_hops_pick(B, t)
        || IREC_StaticInfoExt_note_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_StaticInfoExt_end(B));
}

__flatbuffers_build_table_field(0, flatbuffers_, IREC_SignedExtensions_static_info, IREC_StaticInfoExt, IREC_SignedExtensions)
__flatbuffers_build_struct_field(1, flatbuffers_, IREC_SignedExtensions_hidden_path, IREC_HiddenPathExt, 1, 1, IREC_SignedExtensions)
__flatbuffers_build_struct_field(2, flatbuffers_, IREC_SignedExtensions_irec, IREC_IrecExt, 4, 4, IREC_SignedExtensions)

static inline IREC_SignedExtensions_ref_t IREC_SignedExtensions_create(flatbuffers_builder_t *B __IREC_SignedExtensions_formal_args)
{
    if (IREC_SignedExtensions_start(B)
        || IREC_SignedExtensions_static_info_add(B, v0)
        || IREC_SignedExtensions_irec_add(B, v2)
        || IREC_SignedExtensions_hidden_path_add(B, v1)) {
        return 0;
    }
    return IREC_SignedExtensions_end(B);
}

static IREC_SignedExtensions_ref_t IREC_SignedExtensions_clone(flatbuffers_builder_t *B, IREC_SignedExtensions_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_SignedExtensions_start(B)
        || IREC_SignedExtensions_static_info_pick(B, t)
        || IREC_SignedExtensions_irec_pick(B, t)
        || IREC_SignedExtensions_hidden_path_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_SignedExtensions_end(B));
}


static inline IREC_UnsignedExtensions_ref_t IREC_UnsignedExtensions_create(flatbuffers_builder_t *B __IREC_UnsignedExtensions_formal_args)
{
    if (IREC_UnsignedExtensions_start(B)) {
        return 0;
    }
    return IREC_UnsignedExtensions_end(B);
}

static IREC_UnsignedExtensions_ref_t IREC_UnsignedExtensions_clone(flatbuffers_builder_t *B, IREC_UnsignedExtensions_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_UnsignedExtensions_start(B)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_UnsignedExtensions_end(B));
}

__flatbuffers_build_scalar_field(0, flatbuffers_, IREC_ASEntry_isd_as, flatbuffers_uint64, uint64_t, 8, 8, UINT64_C(0), IREC_ASEntry)
__flatbuffers_build_scalar_field(1, flatbuffers_, IREC_ASEntry_next_isd_as, flatbuffers_uint64, uint64_t, 8, 8, UINT64_C(0), IREC_ASEntry)
__flatbuffers_build_scalar_field(2, flatbuffers_, IREC_ASEntry_he_ingress_mtu, flatbuffers_uint32, uint32_t, 4, 4, UINT32_C(0), IREC_ASEntry)
__flatbuffers_build_struct_field(3, flatbuffers_, IREC_ASEntry_hop_field, IREC_HopField, 24, 8, IREC_ASEntry)
__flatbuffers_build_vector_field(4, flatbuffers_, IREC_ASEntry_peer_entries, IREC_PeerEntry, IREC_PeerEntry_t, IREC_ASEntry)
__flatbuffers_build_scalar_field(5, flatbuffers_, IREC_ASEntry_mtu, flatbuffers_uint32, uint32_t, 4, 4, UINT32_C(0), IREC_ASEntry)
__flatbuffers_build_table_field(6, flatbuffers_, IREC_ASEntry_extensions, IREC_SignedExtensions, IREC_ASEntry)
__flatbuffers_build_table_field(7, flatbuffers_, IREC_ASEntry_unsigned_extensions, IREC_UnsignedExtensions, IREC_ASEntry)

static inline IREC_ASEntry_ref_t IREC_ASEntry_create(flatbuffers_builder_t *B __IREC_ASEntry_formal_args)
{
    if (IREC_ASEntry_start(B)
        || IREC_ASEntry_isd_as_add(B, v0)
        || IREC_ASEntry_next_isd_as_add(B, v1)
        || IREC_ASEntry_hop_field_add(B, v3)
        || IREC_ASEntry_he_ingress_mtu_add(B, v2)
        || IREC_ASEntry_peer_entries_add(B, v4)
        || IREC_ASEntry_mtu_add(B, v5)
        || IREC_ASEntry_extensions_add(B, v6)
        || IREC_ASEntry_unsigned_extensions_add(B, v7)) {
        return 0;
    }
    return IREC_ASEntry_end(B);
}

static IREC_ASEntry_ref_t IREC_ASEntry_clone(flatbuffers_builder_t *B, IREC_ASEntry_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_ASEntry_start(B)
        || IREC_ASEntry_isd_as_pick(B, t)
        || IREC_ASEntry_next_isd_as_pick(B, t)
        || IREC_ASEntry_hop_field_pick(B, t)
        || IREC_ASEntry_he_ingress_mtu_pick(B, t)
        || IREC_ASEntry_peer_entries_pick(B, t)
        || IREC_ASEntry_mtu_pick(B, t)
        || IREC_ASEntry_extensions_pick(B, t)
        || IREC_ASEntry_unsigned_extensions_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_ASEntry_end(B));
}

__flatbuffers_build_vector_field(0, flatbuffers_, IREC_Beacon_segment_info, flatbuffers_int8, int8_t, IREC_Beacon)
__flatbuffers_build_table_vector_field(1, flatbuffers_, IREC_Beacon_as_entries, IREC_ASEntry, IREC_Beacon)
__flatbuffers_build_scalar_field(2, flatbuffers_, IREC_Beacon_inIfId, flatbuffers_uint32, uint32_t, 4, 4, UINT32_C(0), IREC_Beacon)
__flatbuffers_build_scalar_field(3, flatbuffers_, IREC_Beacon_id, flatbuffers_uint32, uint32_t, 4, 4, UINT32_C(0), IREC_Beacon)

static inline IREC_Beacon_ref_t IREC_Beacon_create(flatbuffers_builder_t *B __IREC_Beacon_formal_args)
{
    if (IREC_Beacon_start(B)
        || IREC_Beacon_segment_info_add(B, v0)
        || IREC_Beacon_as_entries_add(B, v1)
        || IREC_Beacon_inIfId_add(B, v2)
        || IREC_Beacon_id_add(B, v3)) {
        return 0;
    }
    return IREC_Beacon_end(B);
}

static IREC_Beacon_ref_t IREC_Beacon_clone(flatbuffers_builder_t *B, IREC_Beacon_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_Beacon_start(B)
        || IREC_Beacon_segment_info_pick(B, t)
        || IREC_Beacon_as_entries_pick(B, t)
        || IREC_Beacon_inIfId_pick(B, t)
        || IREC_Beacon_id_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_Beacon_end(B));
}

__flatbuffers_build_table_vector_field(0, flatbuffers_, IREC_Execution_beacons, IREC_Beacon, IREC_Execution)
__flatbuffers_build_vector_field(1, flatbuffers_, IREC_Execution_algHash, flatbuffers_int8, int8_t, IREC_Execution)
__flatbuffers_build_vector_field(2, flatbuffers_, IREC_Execution_intfs, IREC_Interface, IREC_Interface_t, IREC_Execution)

static inline IREC_Execution_ref_t IREC_Execution_create(flatbuffers_builder_t *B __IREC_Execution_formal_args)
{
    if (IREC_Execution_start(B)
        || IREC_Execution_beacons_add(B, v0)
        || IREC_Execution_algHash_add(B, v1)
        || IREC_Execution_intfs_add(B, v2)) {
        return 0;
    }
    return IREC_Execution_end(B);
}

static IREC_Execution_ref_t IREC_Execution_clone(flatbuffers_builder_t *B, IREC_Execution_table_t t)
{
    __flatbuffers_memoize_begin(B, t);
    if (IREC_Execution_start(B)
        || IREC_Execution_beacons_pick(B, t)
        || IREC_Execution_algHash_pick(B, t)
        || IREC_Execution_intfs_pick(B, t)) {
        return 0;
    }
    __flatbuffers_memoize_end(B, t, IREC_Execution_end(B));
}

#include "flatcc/flatcc_epilogue.h"
#endif /* IREC_BUILDER_H */