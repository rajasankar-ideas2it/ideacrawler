/*************************************************************************
 *
 * Copyright 2018 Ideas2IT Technology Services Private Limited.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 ***********************************************************************/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protofiles/ideacrawler.proto

/*
Package protofiles is a generated protocol buffer package.

It is generated from these files:
	protofiles/ideacrawler.proto

It has these top-level messages:
	Status
	KVP
	DomainOpt
	Subscription
	PageRequest
	PageHTML
*/
package protofiles

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Subscription type
type SubType int32

const (
	// crawler will remember sequence number of each page stored, so we can start back exactly where we left off
	SubType_SEQNUM SubType = 0
	// if we know only the time when we left off,  or if we want to start reading from a certain day's run
	SubType_DATETIME SubType = 1
)

var SubType_name = map[int32]string{
	0: "SEQNUM",
	1: "DATETIME",
}
var SubType_value = map[string]int32{
	"SEQNUM":   0,
	"DATETIME": 1,
}

func (x SubType) String() string {
	return proto.EnumName(SubType_name, int32(x))
}
func (SubType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PageReqType int32

const (
	PageReqType_GET PageReqType = 0
	// Sends a HEAD request to first identify page is text/html before downloading
	// If we are unsure link will send back large gzip file, etc. which we want to
	// avoid.
	PageReqType_HEAD      PageReqType = 1
	PageReqType_BUILTINJS PageReqType = 2
	PageReqType_JSCRIPT   PageReqType = 3
)

var PageReqType_name = map[int32]string{
	0: "GET",
	1: "HEAD",
	2: "BUILTINJS",
	3: "JSCRIPT",
}
var PageReqType_value = map[string]int32{
	"GET":       0,
	"HEAD":      1,
	"BUILTINJS": 2,
	"JSCRIPT":   3,
}

func (x PageReqType) String() string {
	return proto.EnumName(PageReqType_name, int32(x))
}
func (PageReqType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Status struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type KVP struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KVP) Reset()                    { *m = KVP{} }
func (m *KVP) String() string            { return proto.CompactTextString(m) }
func (*KVP) ProtoMessage()               {}
func (*KVP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KVP) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVP) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DomainOpt struct {
	SeedUrl string `protobuf:"bytes,1,opt,name=seedUrl" json:"seedUrl,omitempty"`
	// crawl delay in seconds
	MinDelay int32 `protobuf:"varint,2,opt,name=minDelay" json:"minDelay,omitempty"`
	MaxDelay int32 `protobuf:"varint,3,opt,name=maxDelay" json:"maxDelay,omitempty"`
	// don't follow any pages,  just send back responses for the received URLs.
	NoFollow bool `protobuf:"varint,4,opt,name=noFollow" json:"noFollow,omitempty"`
	// only pages matching reqUrlRegexp will be shipped back to the client.
	// only matching pages will be saved to s3 as well.
	CallbackUrlRegexp string `protobuf:"bytes,5,opt,name=callbackUrlRegexp" json:"callbackUrlRegexp,omitempty"`
	// only pages matching followUrlRegexp will be followed and sublinks added to fetcher.
	FollowUrlRegexp       string `protobuf:"bytes,6,opt,name=followUrlRegexp" json:"followUrlRegexp,omitempty"`
	MaxConcurrentRequests int32  `protobuf:"varint,7,opt,name=maxConcurrentRequests" json:"maxConcurrentRequests,omitempty"`
	// TODO
	Useragent string `protobuf:"bytes,8,opt,name=useragent" json:"useragent,omitempty"`
	Impolite  bool   `protobuf:"varint,9,opt,name=impolite" json:"impolite,omitempty"`
	// TODO
	Depth int32 `protobuf:"varint,10,opt,name=depth" json:"depth,omitempty"`
	// TODO: maybe just remove all scheduling features, immediate jobs only
	Repeat bool `protobuf:"varint,11,opt,name=repeat" json:"repeat,omitempty"`
	// needs min limit of 5mins, ideally 1hour
	Frequency *google_protobuf1.Duration `protobuf:"bytes,12,opt,name=frequency" json:"frequency,omitempty"`
	// time of first run, if this is saturday 10pm, frequency is 2 weeks. ideally atleast 10 mins away.
	// it will continue to run at that time every 2 weeks
	Firstrun *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=firstrun" json:"firstrun,omitempty"`
	// Callback check order -
	//   (1) - callbackUrlRegexp
	//   (2) - callbackXpathMatch
	//   (3) - callbackXpathRegexp
	//  Any one has to match.
	// provide multiple xpaths as keys and expected values as value.  Pages are
	// sent back to client only if all values are found in page.
	CallbackXpathMatch  []*KVP `protobuf:"bytes,14,rep,name=callbackXpathMatch" json:"callbackXpathMatch,omitempty"`
	CallbackXpathRegexp []*KVP `protobuf:"bytes,15,rep,name=callbackXpathRegexp" json:"callbackXpathRegexp,omitempty"`
	//  in seconds, it is the time to wait for a new
	// page, before stopping the job; affects workerIdleTTL of fetchbot.
	// min value is 600, it is also default.
	MaxIdleTime        int32    `protobuf:"varint,16,opt,name=maxIdleTime" json:"maxIdleTime,omitempty"`
	FollowOtherDomains bool     `protobuf:"varint,17,opt,name=followOtherDomains" json:"followOtherDomains,omitempty"`
	KeepDomains        []string `protobuf:"bytes,18,rep,name=keepDomains" json:"keepDomains,omitempty"`
	DropDomains        []string `protobuf:"bytes,19,rep,name=dropDomains" json:"dropDomains,omitempty"`
	DomainDropPriority bool     `protobuf:"varint,20,opt,name=domainDropPriority" json:"domainDropPriority,omitempty"`
	// safe url normalizations happen by default. below is only for a few unsafe ones.
	// for list of safe normalizations: https://github.com/PuerkitoBio/purell/blob/master/purell.go#L59
	// remove index.php, etc,  fragments #section, +FlagsUsuallySafeGreedy from above link
	UnsafeNormalizeURL bool `protobuf:"varint,21,opt,name=unsafeNormalizeURL" json:"unsafeNormalizeURL,omitempty"`
	Login              bool `protobuf:"varint,22,opt,name=login" json:"login,omitempty"`
	// currently not possible, assumes false. uses chrome debugging protocol directly.
	LoginUsingSelenium bool   `protobuf:"varint,23,opt,name=loginUsingSelenium" json:"loginUsingSelenium,omitempty"`
	LoginUrl           string `protobuf:"bytes,24,opt,name=loginUrl" json:"loginUrl,omitempty"`
	// for username, password fields, other form data to send on post request
	LoginPayload []*KVP `protobuf:"bytes,25,rep,name=loginPayload" json:"loginPayload,omitempty"`
	// if there are hidden fields in the page that need to be scraped before login
	LoginParseFields bool `protobuf:"varint,26,opt,name=loginParseFields" json:"loginParseFields,omitempty"`
	// key is key of hidden fields to parse from form, value is the xpath of field to scrape.
	LoginParseXpath []*KVP `protobuf:"bytes,27,rep,name=loginParseXpath" json:"loginParseXpath,omitempty"`
	// to check if login succeeded, provide xpath as key, and expected value as value.
	// for example,  after login, xpath of top right corner,  and username as value.
	// if the xpath is not there of if there is no value match,  then we probably didn't login.
	LoginSuccessCheck *KVP `protobuf:"bytes,28,opt,name=loginSuccessCheck" json:"loginSuccessCheck,omitempty"`
	// checks login state after downloading each page, using check defined in 'loginSuccessCheck'
	CheckLoginAfterEachPage bool `protobuf:"varint,29,opt,name=checkLoginAfterEachPage" json:"checkLoginAfterEachPage,omitempty"`
	// javascript for login in chrome browser.
	LoginJS string `protobuf:"bytes,30,opt,name=loginJS" json:"loginJS,omitempty"`
	// whether to use chrome, location of chrome binary
	Chrome       bool   `protobuf:"varint,31,opt,name=chrome" json:"chrome,omitempty"`
	ChromeBinary string `protobuf:"bytes,32,opt,name=chromeBinary" json:"chromeBinary,omitempty"`
	DomLoadTime  int32  `protobuf:"varint,33,opt,name=domLoadTime" json:"domLoadTime,omitempty"`
	// check if this network interface is still active before every request.
	NetworkIface       string `protobuf:"bytes,34,opt,name=networkIface" json:"networkIface,omitempty"`
	CancelOnDisconnect bool   `protobuf:"varint,35,opt,name=cancelOnDisconnect" json:"cancelOnDisconnect,omitempty"`
	// if true,  sends a HEAD request first ensure content is text/html before sending GET request.
	CheckContent bool `protobuf:"varint,36,opt,name=checkContent" json:"checkContent,omitempty"`
	// if prefetch flag is true, downloads resources like img, css, png, svg, js associated with the actual page to mimic browser behaviour.
	Prefetch bool `protobuf:"varint,37,opt,name=prefetch" json:"prefetch,omitempty"`
}

func (m *DomainOpt) Reset()                    { *m = DomainOpt{} }
func (m *DomainOpt) String() string            { return proto.CompactTextString(m) }
func (*DomainOpt) ProtoMessage()               {}
func (*DomainOpt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DomainOpt) GetSeedUrl() string {
	if m != nil {
		return m.SeedUrl
	}
	return ""
}

func (m *DomainOpt) GetMinDelay() int32 {
	if m != nil {
		return m.MinDelay
	}
	return 0
}

func (m *DomainOpt) GetMaxDelay() int32 {
	if m != nil {
		return m.MaxDelay
	}
	return 0
}

func (m *DomainOpt) GetNoFollow() bool {
	if m != nil {
		return m.NoFollow
	}
	return false
}

func (m *DomainOpt) GetCallbackUrlRegexp() string {
	if m != nil {
		return m.CallbackUrlRegexp
	}
	return ""
}

func (m *DomainOpt) GetFollowUrlRegexp() string {
	if m != nil {
		return m.FollowUrlRegexp
	}
	return ""
}

func (m *DomainOpt) GetMaxConcurrentRequests() int32 {
	if m != nil {
		return m.MaxConcurrentRequests
	}
	return 0
}

func (m *DomainOpt) GetUseragent() string {
	if m != nil {
		return m.Useragent
	}
	return ""
}

func (m *DomainOpt) GetImpolite() bool {
	if m != nil {
		return m.Impolite
	}
	return false
}

func (m *DomainOpt) GetDepth() int32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *DomainOpt) GetRepeat() bool {
	if m != nil {
		return m.Repeat
	}
	return false
}

func (m *DomainOpt) GetFrequency() *google_protobuf1.Duration {
	if m != nil {
		return m.Frequency
	}
	return nil
}

func (m *DomainOpt) GetFirstrun() *google_protobuf.Timestamp {
	if m != nil {
		return m.Firstrun
	}
	return nil
}

func (m *DomainOpt) GetCallbackXpathMatch() []*KVP {
	if m != nil {
		return m.CallbackXpathMatch
	}
	return nil
}

func (m *DomainOpt) GetCallbackXpathRegexp() []*KVP {
	if m != nil {
		return m.CallbackXpathRegexp
	}
	return nil
}

func (m *DomainOpt) GetMaxIdleTime() int32 {
	if m != nil {
		return m.MaxIdleTime
	}
	return 0
}

func (m *DomainOpt) GetFollowOtherDomains() bool {
	if m != nil {
		return m.FollowOtherDomains
	}
	return false
}

func (m *DomainOpt) GetKeepDomains() []string {
	if m != nil {
		return m.KeepDomains
	}
	return nil
}

func (m *DomainOpt) GetDropDomains() []string {
	if m != nil {
		return m.DropDomains
	}
	return nil
}

func (m *DomainOpt) GetDomainDropPriority() bool {
	if m != nil {
		return m.DomainDropPriority
	}
	return false
}

func (m *DomainOpt) GetUnsafeNormalizeURL() bool {
	if m != nil {
		return m.UnsafeNormalizeURL
	}
	return false
}

func (m *DomainOpt) GetLogin() bool {
	if m != nil {
		return m.Login
	}
	return false
}

func (m *DomainOpt) GetLoginUsingSelenium() bool {
	if m != nil {
		return m.LoginUsingSelenium
	}
	return false
}

func (m *DomainOpt) GetLoginUrl() string {
	if m != nil {
		return m.LoginUrl
	}
	return ""
}

func (m *DomainOpt) GetLoginPayload() []*KVP {
	if m != nil {
		return m.LoginPayload
	}
	return nil
}

func (m *DomainOpt) GetLoginParseFields() bool {
	if m != nil {
		return m.LoginParseFields
	}
	return false
}

func (m *DomainOpt) GetLoginParseXpath() []*KVP {
	if m != nil {
		return m.LoginParseXpath
	}
	return nil
}

func (m *DomainOpt) GetLoginSuccessCheck() *KVP {
	if m != nil {
		return m.LoginSuccessCheck
	}
	return nil
}

func (m *DomainOpt) GetCheckLoginAfterEachPage() bool {
	if m != nil {
		return m.CheckLoginAfterEachPage
	}
	return false
}

func (m *DomainOpt) GetLoginJS() string {
	if m != nil {
		return m.LoginJS
	}
	return ""
}

func (m *DomainOpt) GetChrome() bool {
	if m != nil {
		return m.Chrome
	}
	return false
}

func (m *DomainOpt) GetChromeBinary() string {
	if m != nil {
		return m.ChromeBinary
	}
	return ""
}

func (m *DomainOpt) GetDomLoadTime() int32 {
	if m != nil {
		return m.DomLoadTime
	}
	return 0
}

func (m *DomainOpt) GetNetworkIface() string {
	if m != nil {
		return m.NetworkIface
	}
	return ""
}

func (m *DomainOpt) GetCancelOnDisconnect() bool {
	if m != nil {
		return m.CancelOnDisconnect
	}
	return false
}

func (m *DomainOpt) GetCheckContent() bool {
	if m != nil {
		return m.CheckContent
	}
	return false
}

func (m *DomainOpt) GetPrefetch() bool {
	if m != nil {
		return m.Prefetch
	}
	return false
}

type Subscription struct {
	Subcode    string                     `protobuf:"bytes,1,opt,name=subcode" json:"subcode,omitempty"`
	Domainname string                     `protobuf:"bytes,2,opt,name=domainname" json:"domainname,omitempty"`
	Subtype    SubType                    `protobuf:"varint,3,opt,name=subtype,enum=protofiles.SubType" json:"subtype,omitempty"`
	Seqnum     int32                      `protobuf:"varint,4,opt,name=seqnum" json:"seqnum,omitempty"`
	Datetime   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=datetime" json:"datetime,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Subscription) GetSubcode() string {
	if m != nil {
		return m.Subcode
	}
	return ""
}

func (m *Subscription) GetDomainname() string {
	if m != nil {
		return m.Domainname
	}
	return ""
}

func (m *Subscription) GetSubtype() SubType {
	if m != nil {
		return m.Subtype
	}
	return SubType_SEQNUM
}

func (m *Subscription) GetSeqnum() int32 {
	if m != nil {
		return m.Seqnum
	}
	return 0
}

func (m *Subscription) GetDatetime() *google_protobuf.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

type PageRequest struct {
	Sub        *Subscription `protobuf:"bytes,1,opt,name=sub" json:"sub,omitempty"`
	Reqtype    PageReqType   `protobuf:"varint,2,opt,name=reqtype,enum=protofiles.PageReqType" json:"reqtype,omitempty"`
	Url        string        `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Js         string        `protobuf:"bytes,4,opt,name=js" json:"js,omitempty"`
	NoCallback bool          `protobuf:"varint,5,opt,name=noCallback" json:"noCallback,omitempty"`
	MetaStr    string        `protobuf:"bytes,6,opt,name=metaStr" json:"metaStr,omitempty"`
}

func (m *PageRequest) Reset()                    { *m = PageRequest{} }
func (m *PageRequest) String() string            { return proto.CompactTextString(m) }
func (*PageRequest) ProtoMessage()               {}
func (*PageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PageRequest) GetSub() *Subscription {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *PageRequest) GetReqtype() PageReqType {
	if m != nil {
		return m.Reqtype
	}
	return PageReqType_GET
}

func (m *PageRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PageRequest) GetJs() string {
	if m != nil {
		return m.Js
	}
	return ""
}

func (m *PageRequest) GetNoCallback() bool {
	if m != nil {
		return m.NoCallback
	}
	return false
}

func (m *PageRequest) GetMetaStr() string {
	if m != nil {
		return m.MetaStr
	}
	return ""
}

type PageHTML struct {
	Success        bool          `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Error          string        `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Sub            *Subscription `protobuf:"bytes,3,opt,name=sub" json:"sub,omitempty"`
	Url            string        `protobuf:"bytes,4,opt,name=url" json:"url,omitempty"`
	Httpstatuscode int32         `protobuf:"varint,5,opt,name=httpstatuscode" json:"httpstatuscode,omitempty"`
	Content        []byte        `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	MetaStr        string        `protobuf:"bytes,7,opt,name=metaStr" json:"metaStr,omitempty"`
	UrlDepth       int32         `protobuf:"varint,8,opt,name=urlDepth" json:"urlDepth,omitempty"`
}

func (m *PageHTML) Reset()                    { *m = PageHTML{} }
func (m *PageHTML) String() string            { return proto.CompactTextString(m) }
func (*PageHTML) ProtoMessage()               {}
func (*PageHTML) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PageHTML) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PageHTML) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PageHTML) GetSub() *Subscription {
	if m != nil {
		return m.Sub
	}
	return nil
}

func (m *PageHTML) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *PageHTML) GetHttpstatuscode() int32 {
	if m != nil {
		return m.Httpstatuscode
	}
	return 0
}

func (m *PageHTML) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PageHTML) GetMetaStr() string {
	if m != nil {
		return m.MetaStr
	}
	return ""
}

func (m *PageHTML) GetUrlDepth() int32 {
	if m != nil {
		return m.UrlDepth
	}
	return 0
}

func init() {
	proto.RegisterType((*Status)(nil), "protofiles.Status")
	proto.RegisterType((*KVP)(nil), "protofiles.KVP")
	proto.RegisterType((*DomainOpt)(nil), "protofiles.DomainOpt")
	proto.RegisterType((*Subscription)(nil), "protofiles.Subscription")
	proto.RegisterType((*PageRequest)(nil), "protofiles.PageRequest")
	proto.RegisterType((*PageHTML)(nil), "protofiles.PageHTML")
	proto.RegisterEnum("protofiles.SubType", SubType_name, SubType_value)
	proto.RegisterEnum("protofiles.PageReqType", PageReqType_name, PageReqType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IdeaCrawler service

type IdeaCrawlerClient interface {
	AddDomainAndListen(ctx context.Context, in *DomainOpt, opts ...grpc.CallOption) (IdeaCrawler_AddDomainAndListenClient, error)
	AddPages(ctx context.Context, opts ...grpc.CallOption) (IdeaCrawler_AddPagesClient, error)
	CancelJob(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Status, error)
}

type ideaCrawlerClient struct {
	cc *grpc.ClientConn
}

func NewIdeaCrawlerClient(cc *grpc.ClientConn) IdeaCrawlerClient {
	return &ideaCrawlerClient{cc}
}

func (c *ideaCrawlerClient) AddDomainAndListen(ctx context.Context, in *DomainOpt, opts ...grpc.CallOption) (IdeaCrawler_AddDomainAndListenClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_IdeaCrawler_serviceDesc.Streams[0], c.cc, "/protofiles.IdeaCrawler/AddDomainAndListen", opts...)
	if err != nil {
		return nil, err
	}
	x := &ideaCrawlerAddDomainAndListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IdeaCrawler_AddDomainAndListenClient interface {
	Recv() (*PageHTML, error)
	grpc.ClientStream
}

type ideaCrawlerAddDomainAndListenClient struct {
	grpc.ClientStream
}

func (x *ideaCrawlerAddDomainAndListenClient) Recv() (*PageHTML, error) {
	m := new(PageHTML)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ideaCrawlerClient) AddPages(ctx context.Context, opts ...grpc.CallOption) (IdeaCrawler_AddPagesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_IdeaCrawler_serviceDesc.Streams[1], c.cc, "/protofiles.IdeaCrawler/AddPages", opts...)
	if err != nil {
		return nil, err
	}
	x := &ideaCrawlerAddPagesClient{stream}
	return x, nil
}

type IdeaCrawler_AddPagesClient interface {
	Send(*PageRequest) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type ideaCrawlerAddPagesClient struct {
	grpc.ClientStream
}

func (x *ideaCrawlerAddPagesClient) Send(m *PageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ideaCrawlerAddPagesClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ideaCrawlerClient) CancelJob(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/protofiles.IdeaCrawler/CancelJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IdeaCrawler service

type IdeaCrawlerServer interface {
	AddDomainAndListen(*DomainOpt, IdeaCrawler_AddDomainAndListenServer) error
	AddPages(IdeaCrawler_AddPagesServer) error
	CancelJob(context.Context, *Subscription) (*Status, error)
}

func RegisterIdeaCrawlerServer(s *grpc.Server, srv IdeaCrawlerServer) {
	s.RegisterService(&_IdeaCrawler_serviceDesc, srv)
}

func _IdeaCrawler_AddDomainAndListen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DomainOpt)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdeaCrawlerServer).AddDomainAndListen(m, &ideaCrawlerAddDomainAndListenServer{stream})
}

type IdeaCrawler_AddDomainAndListenServer interface {
	Send(*PageHTML) error
	grpc.ServerStream
}

type ideaCrawlerAddDomainAndListenServer struct {
	grpc.ServerStream
}

func (x *ideaCrawlerAddDomainAndListenServer) Send(m *PageHTML) error {
	return x.ServerStream.SendMsg(m)
}

func _IdeaCrawler_AddPages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IdeaCrawlerServer).AddPages(&ideaCrawlerAddPagesServer{stream})
}

type IdeaCrawler_AddPagesServer interface {
	SendAndClose(*Status) error
	Recv() (*PageRequest, error)
	grpc.ServerStream
}

type ideaCrawlerAddPagesServer struct {
	grpc.ServerStream
}

func (x *ideaCrawlerAddPagesServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ideaCrawlerAddPagesServer) Recv() (*PageRequest, error) {
	m := new(PageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _IdeaCrawler_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdeaCrawlerServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.IdeaCrawler/CancelJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdeaCrawlerServer).CancelJob(ctx, req.(*Subscription))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdeaCrawler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.IdeaCrawler",
	HandlerType: (*IdeaCrawlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelJob",
			Handler:    _IdeaCrawler_CancelJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddDomainAndListen",
			Handler:       _IdeaCrawler_AddDomainAndListen_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddPages",
			Handler:       _IdeaCrawler_AddPages_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protofiles/ideacrawler.proto",
}

func init() { proto.RegisterFile("protofiles/ideacrawler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x6f, 0x73, 0x1b, 0xb5,
	0x13, 0x8e, 0xe3, 0x26, 0xb6, 0xe5, 0x34, 0x71, 0xd5, 0x7f, 0xaa, 0x7f, 0xfd, 0xb5, 0xc6, 0x05,
	0xc6, 0x93, 0xa1, 0x2e, 0xa4, 0x0c, 0x94, 0x01, 0x86, 0x71, 0x6d, 0x97, 0x3a, 0x4d, 0xda, 0x70,
	0x76, 0x3a, 0xbc, 0x95, 0xef, 0xd6, 0xf6, 0x35, 0x77, 0xd2, 0x55, 0xa7, 0x23, 0x31, 0x5f, 0x10,
	0x5e, 0xf1, 0x0d, 0xf8, 0x0e, 0x7c, 0x04, 0x46, 0xab, 0x3b, 0xff, 0xef, 0x00, 0xef, 0xb4, 0xcf,
	0x3e, 0x5a, 0xed, 0x6a, 0x9f, 0x5b, 0x1d, 0xb9, 0x1f, 0x29, 0xa9, 0xe5, 0xc8, 0x0f, 0x20, 0x7e,
	0xe2, 0x7b, 0xc0, 0x5d, 0xc5, 0x2f, 0x03, 0x50, 0x4d, 0x84, 0x29, 0x99, 0x7b, 0xab, 0x0f, 0xc7,
	0x52, 0x8e, 0x03, 0x78, 0x82, 0xd0, 0x30, 0x19, 0x3d, 0xd1, 0x7e, 0x08, 0xb1, 0xe6, 0x61, 0x64,
	0xc9, 0xd5, 0x07, 0xab, 0x04, 0x2f, 0x51, 0x5c, 0xfb, 0x52, 0x58, 0x7f, 0xfd, 0x19, 0xd9, 0xed,
	0x6b, 0xae, 0x93, 0x98, 0x32, 0x52, 0x88, 0x13, 0xd7, 0x85, 0x38, 0x66, 0xb9, 0x5a, 0xae, 0x51,
	0x74, 0x32, 0x93, 0xde, 0x22, 0x3b, 0xa0, 0x94, 0x54, 0x6c, 0xbb, 0x96, 0x6b, 0x94, 0x1c, 0x6b,
	0xd4, 0x1f, 0x93, 0xfc, 0xab, 0xb7, 0x67, 0xb4, 0x42, 0xf2, 0x17, 0x30, 0xc5, 0x2d, 0x25, 0xc7,
	0x2c, 0x0d, 0xfd, 0x17, 0x1e, 0x24, 0x90, 0xd1, 0xd1, 0xa8, 0xff, 0x59, 0x26, 0xa5, 0x8e, 0x0c,
	0xb9, 0x2f, 0xde, 0x44, 0x1a, 0x0f, 0x03, 0xf0, 0xce, 0x55, 0x90, 0xee, 0xcc, 0x4c, 0x5a, 0x25,
	0xc5, 0xd0, 0x17, 0x1d, 0x08, 0xf8, 0x14, 0x03, 0xec, 0x38, 0x33, 0x1b, 0x7d, 0xfc, 0xca, 0xfa,
	0xf2, 0xa9, 0x2f, 0xb5, 0x8d, 0x4f, 0xc8, 0x17, 0x32, 0x08, 0xe4, 0x25, 0xbb, 0x86, 0xf9, 0xcf,
	0x6c, 0xfa, 0x19, 0xb9, 0xe1, 0xf2, 0x20, 0x18, 0x72, 0xf7, 0xe2, 0x5c, 0x05, 0x0e, 0x8c, 0xe1,
	0x2a, 0x62, 0x3b, 0x78, 0xee, 0xba, 0x83, 0x36, 0xc8, 0xc1, 0x08, 0xf7, 0xcd, 0xb9, 0xbb, 0xc8,
	0x5d, 0x85, 0xe9, 0x97, 0xe4, 0x76, 0xc8, 0xaf, 0xda, 0x52, 0xb8, 0x89, 0x52, 0x20, 0xb4, 0x03,
	0xef, 0x13, 0x88, 0x75, 0xcc, 0x0a, 0x98, 0xdc, 0x66, 0x27, 0xbd, 0x4f, 0x4a, 0x49, 0x0c, 0x8a,
	0x8f, 0x41, 0x68, 0x56, 0xc4, 0xc8, 0x73, 0xc0, 0xd4, 0xe1, 0x87, 0x91, 0x0c, 0x7c, 0x0d, 0xac,
	0x64, 0xeb, 0xc8, 0x6c, 0x73, 0xb3, 0x1e, 0x44, 0x7a, 0xc2, 0x08, 0xc6, 0xb7, 0x06, 0xbd, 0x43,
	0x76, 0x15, 0x44, 0xc0, 0x35, 0x2b, 0x23, 0x3f, 0xb5, 0xe8, 0xd7, 0xa4, 0x34, 0x52, 0xe6, 0x50,
	0xe1, 0x4e, 0xd9, 0x5e, 0x2d, 0xd7, 0x28, 0x1f, 0xdd, 0x6b, 0x5a, 0x39, 0x34, 0x33, 0x39, 0x34,
	0x3b, 0xa9, 0x1c, 0x9c, 0x39, 0x97, 0x7e, 0x45, 0x8a, 0x23, 0x5f, 0xc5, 0x5a, 0x25, 0x82, 0x5d,
	0xc7, 0x7d, 0xd5, 0xb5, 0x7d, 0x83, 0x4c, 0x67, 0xce, 0x8c, 0x4b, 0x7f, 0x20, 0x34, 0xbb, 0xcd,
	0x9f, 0x23, 0xae, 0x27, 0xa7, 0x5c, 0xbb, 0x13, 0xb6, 0x5f, 0xcb, 0x37, 0xca, 0x47, 0x07, 0xcd,
	0xb9, 0x6a, 0x9b, 0xaf, 0xde, 0x9e, 0x39, 0x1b, 0xa8, 0xb4, 0x45, 0x6e, 0x2e, 0xa1, 0xe9, 0xed,
	0x1f, 0x6c, 0x8e, 0xb0, 0x89, 0x4b, 0x6b, 0xa4, 0x1c, 0xf2, 0xab, 0x9e, 0x17, 0x80, 0xc9, 0x90,
	0x55, 0xf0, 0xa2, 0x16, 0x21, 0xda, 0x24, 0xd4, 0xf6, 0xf1, 0x8d, 0x9e, 0x80, 0xb2, 0x92, 0x8c,
	0xd9, 0x0d, 0xbc, 0xba, 0x0d, 0x1e, 0x13, 0xf1, 0x02, 0x20, 0xca, 0x88, 0xb4, 0x96, 0x6f, 0x94,
	0x9c, 0x45, 0xc8, 0x30, 0x3c, 0x25, 0x67, 0x8c, 0x9b, 0x96, 0xb1, 0x00, 0x99, 0x33, 0x3d, 0x5c,
	0x76, 0x94, 0x8c, 0xce, 0x94, 0x2f, 0x95, 0xaf, 0xa7, 0xec, 0x96, 0x3d, 0x73, 0xdd, 0x63, 0xf8,
	0x89, 0x88, 0xf9, 0x08, 0x5e, 0x4b, 0x15, 0xf2, 0xc0, 0xff, 0x15, 0xce, 0x9d, 0x13, 0x76, 0xdb,
	0xf2, 0xd7, 0x3d, 0x46, 0x18, 0x81, 0x1c, 0xfb, 0x82, 0xdd, 0x41, 0x8a, 0x35, 0x4c, 0x14, 0x5c,
	0x9c, 0xc7, 0xbe, 0x18, 0xf7, 0x21, 0x00, 0xe1, 0x27, 0x21, 0xbb, 0x6b, 0xa3, 0xac, 0x7b, 0x8c,
	0xf4, 0x2c, 0xaa, 0x02, 0xc6, 0x50, 0x97, 0x33, 0x9b, 0x3e, 0x25, 0x7b, 0xb8, 0x3e, 0xe3, 0xd3,
	0x40, 0x72, 0x8f, 0xdd, 0xdb, 0xdc, 0x93, 0x25, 0x12, 0x3d, 0x24, 0x95, 0xd4, 0x56, 0x31, 0xbc,
	0xf0, 0x21, 0xf0, 0x62, 0x56, 0xc5, 0xe3, 0xd7, 0x70, 0xfa, 0x0d, 0x39, 0x98, 0x63, 0xd8, 0x51,
	0xf6, 0xbf, 0xcd, 0x67, 0xac, 0xf2, 0xe8, 0xf7, 0xe4, 0x06, 0x42, 0x7d, 0x3b, 0xaf, 0xda, 0x13,
	0x70, 0x2f, 0xd8, 0x7d, 0x14, 0xee, 0xda, 0xe6, 0x75, 0x26, 0x7d, 0x46, 0xee, 0xba, 0x66, 0x71,
	0x62, 0x3c, 0xad, 0x91, 0x06, 0xd5, 0xe5, 0xee, 0xe4, 0x8c, 0x8f, 0x81, 0xfd, 0x1f, 0x93, 0xfd,
	0x90, 0xdb, 0x4c, 0x31, 0x0c, 0x77, 0xdc, 0x67, 0x0f, 0xec, 0x14, 0x4b, 0x4d, 0xf3, 0x4d, 0xba,
	0x13, 0x25, 0x43, 0x60, 0x0f, 0xed, 0x37, 0x69, 0x2d, 0x5a, 0x27, 0x7b, 0x76, 0xf5, 0xdc, 0x17,
	0x5c, 0x4d, 0x59, 0x0d, 0xb7, 0x2d, 0x61, 0x28, 0x27, 0x19, 0x9e, 0x48, 0xee, 0xa1, 0x84, 0x3f,
	0xb2, 0x12, 0x5e, 0x80, 0x4c, 0x14, 0x01, 0xfa, 0x52, 0xaa, 0x8b, 0xde, 0x88, 0xbb, 0xc0, 0xea,
	0x36, 0xca, 0x22, 0x66, 0x9a, 0xef, 0x72, 0xe1, 0x42, 0xf0, 0x46, 0x74, 0xfc, 0xd8, 0x95, 0x42,
	0x80, 0xab, 0xd9, 0x23, 0xdb, 0xfc, 0x75, 0x8f, 0xcd, 0x0c, 0xdc, 0x8b, 0xb6, 0x14, 0xda, 0x0c,
	0xa6, 0x8f, 0x91, 0xb9, 0x84, 0x19, 0x81, 0x44, 0x0a, 0x46, 0x60, 0x3e, 0xeb, 0x4f, 0xec, 0x6c,
	0xca, 0xec, 0xfa, 0x6f, 0x39, 0xb2, 0xd7, 0x4f, 0x86, 0xb1, 0xab, 0xfc, 0xc8, 0x0c, 0x14, 0xfb,
	0x9e, 0x0c, 0x5d, 0xe9, 0xc1, 0x6c, 0xc4, 0x5b, 0x93, 0x3e, 0x20, 0xc4, 0x6a, 0x5e, 0xf0, 0x30,
	0x7b, 0x25, 0x16, 0x10, 0xfa, 0x18, 0x77, 0xea, 0x69, 0x04, 0x38, 0xe5, 0xf7, 0x8f, 0x6e, 0x2e,
	0x76, 0xb1, 0x9f, 0x0c, 0x07, 0xd3, 0x08, 0x9c, 0x8c, 0x63, 0xee, 0x3a, 0x86, 0xf7, 0x22, 0x09,
	0x71, 0xee, 0xef, 0x38, 0xa9, 0x65, 0xc6, 0x98, 0xc7, 0x35, 0x98, 0x17, 0x11, 0x87, 0xfd, 0x3f,
	0x8c, 0xb1, 0x8c, 0x5b, 0xff, 0x3d, 0x47, 0xca, 0xa6, 0xbd, 0xe9, 0xc0, 0xa6, 0x87, 0x24, 0x1f,
	0x27, 0x43, 0x2c, 0xa2, 0x7c, 0xc4, 0x56, 0x52, 0x99, 0xd5, 0xeb, 0x18, 0x12, 0xfd, 0x82, 0x14,
	0x14, 0xbc, 0xc7, 0xd4, 0xb7, 0x31, 0xf5, 0xbb, 0x8b, 0xfc, 0x34, 0xaa, 0x4d, 0x3f, 0xe5, 0x99,
	0x07, 0x34, 0x51, 0x01, 0x56, 0x5a, 0x72, 0xcc, 0x92, 0xee, 0x93, 0xed, 0x77, 0x31, 0x16, 0x53,
	0x72, 0xb6, 0xdf, 0xc5, 0xe6, 0xbe, 0x84, 0x6c, 0xa7, 0xc3, 0x0e, 0x4b, 0x29, 0x3a, 0x0b, 0x88,
	0xb9, 0xe9, 0x10, 0x34, 0xef, 0x6b, 0x95, 0x3e, 0x54, 0x99, 0x59, 0xff, 0x2b, 0x47, 0x8a, 0xe6,
	0xd0, 0x97, 0x83, 0xd3, 0x93, 0xff, 0xfa, 0xc0, 0x67, 0x75, 0xe7, 0xff, 0x4d, 0xdd, 0x69, 0x11,
	0xd7, 0xe6, 0x45, 0x7c, 0x4a, 0xf6, 0x27, 0x5a, 0x47, 0x31, 0xfe, 0x5c, 0xa0, 0x0a, 0x76, 0xb0,
	0x3b, 0x2b, 0xa8, 0xc9, 0xca, 0x4d, 0x25, 0x67, 0x92, 0xdf, 0x73, 0x32, 0x73, 0xb1, 0xac, 0xc2,
	0x52, 0x59, 0x46, 0x87, 0x89, 0x0a, 0x3a, 0xf8, 0x14, 0x16, 0xed, 0x7f, 0x40, 0x66, 0x1f, 0x3e,
	0x22, 0x85, 0x54, 0x21, 0x94, 0x90, 0xdd, 0x7e, 0xf7, 0xa7, 0xd7, 0xe7, 0xa7, 0x95, 0x2d, 0xba,
	0x47, 0x8a, 0x9d, 0xd6, 0xa0, 0x3b, 0xe8, 0x9d, 0x76, 0x2b, 0xb9, 0xc3, 0xef, 0x66, 0x1d, 0x46,
	0x62, 0x81, 0xe4, 0x7f, 0xec, 0x0e, 0x2a, 0x5b, 0xb4, 0x48, 0xae, 0xbd, 0xec, 0xb6, 0x3a, 0x95,
	0x1c, 0xbd, 0x4e, 0x4a, 0xcf, 0xcf, 0x7b, 0x27, 0x83, 0xde, 0xeb, 0xe3, 0x7e, 0x65, 0x9b, 0x96,
	0x49, 0xe1, 0xb8, 0xdf, 0x76, 0x7a, 0x67, 0x83, 0x4a, 0xfe, 0xe8, 0x8f, 0x1c, 0x29, 0xf7, 0x3c,
	0xe0, 0x6d, 0xfb, 0x5b, 0x46, 0xbb, 0x84, 0xb6, 0x3c, 0xcf, 0xce, 0xfa, 0x96, 0xf0, 0x4e, 0xfc,
	0x58, 0x83, 0xa0, 0xb7, 0x17, 0x6f, 0x6c, 0xf6, 0xe7, 0x53, 0xbd, 0xb5, 0x2a, 0x08, 0xd3, 0x9b,
	0xfa, 0xd6, 0xe7, 0x39, 0xfa, 0x2d, 0x29, 0xb6, 0x3c, 0xcf, 0x40, 0x31, 0xdd, 0x24, 0x1b, 0x23,
	0xc6, 0x2a, 0x5d, 0xea, 0x03, 0x5e, 0x63, 0x7d, 0xab, 0x61, 0x36, 0x97, 0xda, 0xf8, 0x51, 0x1f,
	0xcb, 0x21, 0xfd, 0x60, 0xb3, 0x36, 0x6f, 0x1f, 0xee, 0x22, 0xf8, 0xf4, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x50, 0x91, 0x1d, 0x45, 0x78, 0x0a, 0x00, 0x00,
}
