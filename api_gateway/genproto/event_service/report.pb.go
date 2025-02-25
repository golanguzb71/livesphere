// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: report.proto

package event_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IFTAreport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitId         string  `protobuf:"bytes,1,opt,name=unit_id,json=unitId,proto3" json:"unit_id"`
	Miles          float32 `protobuf:"fixed32,2,opt,name=miles,proto3" json:"miles"`
	Month          string  `protobuf:"bytes,3,opt,name=month,proto3" json:"month"`
	State          string  `protobuf:"bytes,4,opt,name=state,proto3" json:"state"`
	Unit           *Object `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit"`
	Year           int32   `protobuf:"varint,6,opt,name=year,proto3" json:"year"`
	CompanyName    string  `protobuf:"bytes,7,opt,name=company_name,json=companyName,proto3" json:"company_name"`
	CompanyAddress string  `protobuf:"bytes,8,opt,name=company_address,json=companyAddress,proto3" json:"company_address"`
}

func (x *IFTAreport) Reset() {
	*x = IFTAreport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFTAreport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFTAreport) ProtoMessage() {}

func (x *IFTAreport) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFTAreport.ProtoReflect.Descriptor instead.
func (*IFTAreport) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{0}
}

func (x *IFTAreport) GetUnitId() string {
	if x != nil {
		return x.UnitId
	}
	return ""
}

func (x *IFTAreport) GetMiles() float32 {
	if x != nil {
		return x.Miles
	}
	return 0
}

func (x *IFTAreport) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *IFTAreport) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *IFTAreport) GetUnit() *Object {
	if x != nil {
		return x.Unit
	}
	return nil
}

func (x *IFTAreport) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *IFTAreport) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *IFTAreport) GetCompanyAddress() string {
	if x != nil {
		return x.CompanyAddress
	}
	return ""
}

type IFTAreportList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count         int32         `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Items         []*IFTAreport `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
	NonTotalMiles float32       `protobuf:"fixed32,3,opt,name=non_total_miles,json=nonTotalMiles,proto3" json:"non_total_miles"`
	TotalMiles    float32       `protobuf:"fixed32,4,opt,name=total_miles,json=totalMiles,proto3" json:"total_miles"`
}

func (x *IFTAreportList) Reset() {
	*x = IFTAreportList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFTAreportList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFTAreportList) ProtoMessage() {}

func (x *IFTAreportList) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFTAreportList.ProtoReflect.Descriptor instead.
func (*IFTAreportList) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{1}
}

func (x *IFTAreportList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *IFTAreportList) GetItems() []*IFTAreport {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *IFTAreportList) GetNonTotalMiles() float32 {
	if x != nil {
		return x.NonTotalMiles
	}
	return 0
}

func (x *IFTAreportList) GetTotalMiles() float32 {
	if x != nil {
		return x.TotalMiles
	}
	return 0
}

type IFTAReportByState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State          string  `protobuf:"bytes,1,opt,name=state,proto3" json:"state"`
	Total          float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total"`
	CompanyName    string  `protobuf:"bytes,3,opt,name=company_name,json=companyName,proto3" json:"company_name"`
	CompanyAddress string  `protobuf:"bytes,4,opt,name=company_address,json=companyAddress,proto3" json:"company_address"`
}

func (x *IFTAReportByState) Reset() {
	*x = IFTAReportByState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFTAReportByState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFTAReportByState) ProtoMessage() {}

func (x *IFTAReportByState) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFTAReportByState.ProtoReflect.Descriptor instead.
func (*IFTAReportByState) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{2}
}

func (x *IFTAReportByState) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *IFTAReportByState) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *IFTAReportByState) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *IFTAReportByState) GetCompanyAddress() string {
	if x != nil {
		return x.CompanyAddress
	}
	return ""
}

type IFTAReportByStateList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32                `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Items []*IFTAReportByState `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *IFTAReportByStateList) Reset() {
	*x = IFTAReportByStateList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFTAReportByStateList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFTAReportByStateList) ProtoMessage() {}

func (x *IFTAReportByStateList) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFTAReportByStateList.ProtoReflect.Descriptor instead.
func (*IFTAReportByStateList) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{3}
}

func (x *IFTAReportByStateList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *IFTAReportByStateList) GetItems() []*IFTAReportByState {
	if x != nil {
		return x.Items
	}
	return nil
}

type IFTAreportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quarter string   `protobuf:"bytes,2,opt,name=quarter,proto3" json:"quarter"`
	States  []string `protobuf:"bytes,3,rep,name=states,proto3" json:"states"`
	Units   []string `protobuf:"bytes,4,rep,name=units,proto3" json:"units"`
	Year    string   `protobuf:"bytes,5,opt,name=year,proto3" json:"year"`
	Month   int32    `protobuf:"varint,6,opt,name=month,proto3" json:"month"`
}

func (x *IFTAreportRequest) Reset() {
	*x = IFTAreportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFTAreportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFTAreportRequest) ProtoMessage() {}

func (x *IFTAreportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFTAreportRequest.ProtoReflect.Descriptor instead.
func (*IFTAreportRequest) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{4}
}

func (x *IFTAreportRequest) GetQuarter() string {
	if x != nil {
		return x.Quarter
	}
	return ""
}

func (x *IFTAreportRequest) GetStates() []string {
	if x != nil {
		return x.States
	}
	return nil
}

func (x *IFTAreportRequest) GetUnits() []string {
	if x != nil {
		return x.Units
	}
	return nil
}

func (x *IFTAreportRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *IFTAreportRequest) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

type ActivityReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance      float32 `protobuf:"fixed32,1,opt,name=distance,proto3" json:"distance"`
	Driver        *Object `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver"`
	DrivingHours  float32 `protobuf:"fixed32,3,opt,name=driving_hours,json=drivingHours,proto3" json:"driving_hours"`
	OdometerEnd   float32 `protobuf:"fixed32,4,opt,name=odometer_end,json=odometerEnd,proto3" json:"odometer_end"`
	OdometerStart float32 `protobuf:"fixed32,5,opt,name=odometer_start,json=odometerStart,proto3" json:"odometer_start"`
	Unit          *Object `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit"`
}

func (x *ActivityReport) Reset() {
	*x = ActivityReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityReport) ProtoMessage() {}

func (x *ActivityReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityReport.ProtoReflect.Descriptor instead.
func (*ActivityReport) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{5}
}

func (x *ActivityReport) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *ActivityReport) GetDriver() *Object {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *ActivityReport) GetDrivingHours() float32 {
	if x != nil {
		return x.DrivingHours
	}
	return 0
}

func (x *ActivityReport) GetOdometerEnd() float32 {
	if x != nil {
		return x.OdometerEnd
	}
	return 0
}

func (x *ActivityReport) GetOdometerStart() float32 {
	if x != nil {
		return x.OdometerStart
	}
	return 0
}

func (x *ActivityReport) GetUnit() *Object {
	if x != nil {
		return x.Unit
	}
	return nil
}

type ActivityReportList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32             `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Items []*ActivityReport `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *ActivityReportList) Reset() {
	*x = ActivityReportList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityReportList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityReportList) ProtoMessage() {}

func (x *ActivityReportList) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityReportList.ProtoReflect.Descriptor instead.
func (*ActivityReportList) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{6}
}

func (x *ActivityReportList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ActivityReportList) GetItems() []*ActivityReport {
	if x != nil {
		return x.Items
	}
	return nil
}

type DvirReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date              string       `protobuf:"bytes,1,opt,name=date,proto3" json:"date"`
	Defects           []*Defects   `protobuf:"bytes,2,rep,name=defects,proto3" json:"defects"`
	DriverSignature   string       `protobuf:"bytes,3,opt,name=driver_signature,json=driverSignature,proto3" json:"driver_signature"`
	MechanicSignature string       `protobuf:"bytes,4,opt,name=mechanic_signature,json=mechanicSignature,proto3" json:"mechanic_signature"`
	Status            string       `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`
	Id                string       `protobuf:"bytes,6,opt,name=id,proto3" json:"id"`
	UnitId            string       `protobuf:"bytes,7,opt,name=unit_id,json=unitId,proto3" json:"unit_id"`
	Driver            *Object      `protobuf:"bytes,8,opt,name=driver,proto3" json:"driver"`
	Address           string       `protobuf:"bytes,9,opt,name=address,proto3" json:"address"`
	Gps               *Coordinates `protobuf:"bytes,10,opt,name=gps,proto3" json:"gps"`
	Odometr           float32      `protobuf:"fixed32,11,opt,name=odometr,proto3" json:"odometr"`
	Trailer           string       `protobuf:"bytes,12,opt,name=trailer,proto3" json:"trailer"`
	CompanyId         string       `protobuf:"bytes,13,opt,name=company_id,json=companyId,proto3" json:"company_id"`
	Link              string       `protobuf:"bytes,14,opt,name=link,proto3" json:"link"`
}

func (x *DvirReport) Reset() {
	*x = DvirReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DvirReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DvirReport) ProtoMessage() {}

func (x *DvirReport) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DvirReport.ProtoReflect.Descriptor instead.
func (*DvirReport) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{7}
}

func (x *DvirReport) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *DvirReport) GetDefects() []*Defects {
	if x != nil {
		return x.Defects
	}
	return nil
}

func (x *DvirReport) GetDriverSignature() string {
	if x != nil {
		return x.DriverSignature
	}
	return ""
}

func (x *DvirReport) GetMechanicSignature() string {
	if x != nil {
		return x.MechanicSignature
	}
	return ""
}

func (x *DvirReport) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DvirReport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DvirReport) GetUnitId() string {
	if x != nil {
		return x.UnitId
	}
	return ""
}

func (x *DvirReport) GetDriver() *Object {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *DvirReport) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DvirReport) GetGps() *Coordinates {
	if x != nil {
		return x.Gps
	}
	return nil
}

func (x *DvirReport) GetOdometr() float32 {
	if x != nil {
		return x.Odometr
	}
	return 0
}

func (x *DvirReport) GetTrailer() string {
	if x != nil {
		return x.Trailer
	}
	return ""
}

func (x *DvirReport) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *DvirReport) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type DvirReportList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32         `protobuf:"varint,1,opt,name=count,proto3" json:"count"`
	Items []*DvirReport `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
}

func (x *DvirReportList) Reset() {
	*x = DvirReportList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DvirReportList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DvirReportList) ProtoMessage() {}

func (x *DvirReportList) ProtoReflect() protoreflect.Message {
	mi := &file_report_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DvirReportList.ProtoReflect.Descriptor instead.
func (*DvirReportList) Descriptor() ([]byte, []int) {
	return file_report_proto_rawDescGZIP(), []int{8}
}

func (x *DvirReportList) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DvirReportList) GetItems() []*DvirReport {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_report_proto protoreflect.FileDescriptor

var file_report_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x12, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x65, 0x6c, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x0a, 0x49, 0x46, 0x54, 0x41, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x49, 0x46, 0x54, 0x41, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x46, 0x54, 0x41, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x6e, 0x6f, 0x6e, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x49, 0x46,
	0x54, 0x41, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x65, 0x0a, 0x15, 0x49, 0x46, 0x54, 0x41, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x46, 0x54, 0x41, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x42, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x85,
	0x01, 0x0a, 0x11, 0x49, 0x46, 0x54, 0x41, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x61, 0x72, 0x74, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x22, 0xf5, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x64, 0x72, 0x69,
	0x76, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x64, 0x6f,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0b, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x5f,
	0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0xcb, 0x03, 0x0a, 0x0a, 0x44, 0x76, 0x69, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x66, 0x65, 0x63, 0x74, 0x73, 0x52, 0x07, 0x64, 0x65, 0x66,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x63, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x03, 0x67, 0x70, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x03, 0x67, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6f, 0x64, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x57, 0x0a,
	0x0e, 0x44, 0x76, 0x69, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x76, 0x69, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_proto_rawDescOnce sync.Once
	file_report_proto_rawDescData = file_report_proto_rawDesc
)

func file_report_proto_rawDescGZIP() []byte {
	file_report_proto_rawDescOnce.Do(func() {
		file_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_proto_rawDescData)
	})
	return file_report_proto_rawDescData
}

var file_report_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_report_proto_goTypes = []any{
	(*IFTAreport)(nil),            // 0: event_service.IFTAreport
	(*IFTAreportList)(nil),        // 1: event_service.IFTAreportList
	(*IFTAReportByState)(nil),     // 2: event_service.IFTAReportByState
	(*IFTAReportByStateList)(nil), // 3: event_service.IFTAReportByStateList
	(*IFTAreportRequest)(nil),     // 4: event_service.IFTAreportRequest
	(*ActivityReport)(nil),        // 5: event_service.ActivityReport
	(*ActivityReportList)(nil),    // 6: event_service.ActivityReportList
	(*DvirReport)(nil),            // 7: event_service.DvirReport
	(*DvirReportList)(nil),        // 8: event_service.DvirReportList
	(*Object)(nil),                // 9: event_service.Object
	(*Defects)(nil),               // 10: event_service.Defects
	(*Coordinates)(nil),           // 11: event_service.Coordinates
}
var file_report_proto_depIdxs = []int32{
	9,  // 0: event_service.IFTAreport.unit:type_name -> event_service.Object
	0,  // 1: event_service.IFTAreportList.items:type_name -> event_service.IFTAreport
	2,  // 2: event_service.IFTAReportByStateList.items:type_name -> event_service.IFTAReportByState
	9,  // 3: event_service.ActivityReport.driver:type_name -> event_service.Object
	9,  // 4: event_service.ActivityReport.unit:type_name -> event_service.Object
	5,  // 5: event_service.ActivityReportList.items:type_name -> event_service.ActivityReport
	10, // 6: event_service.DvirReport.defects:type_name -> event_service.Defects
	9,  // 7: event_service.DvirReport.driver:type_name -> event_service.Object
	11, // 8: event_service.DvirReport.gps:type_name -> event_service.Coordinates
	7,  // 9: event_service.DvirReportList.items:type_name -> event_service.DvirReport
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_report_proto_init() }
func file_report_proto_init() {
	if File_report_proto != nil {
		return
	}
	file_event_models_proto_init()
	file_eld_logs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_report_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IFTAreport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IFTAreportList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IFTAReportByState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IFTAReportByStateList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IFTAreportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityReportList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DvirReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_report_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DvirReportList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_report_proto_goTypes,
		DependencyIndexes: file_report_proto_depIdxs,
		MessageInfos:      file_report_proto_msgTypes,
	}.Build()
	File_report_proto = out.File
	file_report_proto_rawDesc = nil
	file_report_proto_goTypes = nil
	file_report_proto_depIdxs = nil
}
