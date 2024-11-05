package announcement

import "github.com/kysion/base-library/utility/enum"

// StateEnum 状态： 1草稿、2待发布、4已发布、8已过期、16已撤销、
type StateEnum enum.IEnumCode[int]

type state struct {
	Draft       StateEnum
	WaitRelease StateEnum
	Published   StateEnum
	Expired     StateEnum
	Revoked     StateEnum
}

var State = state{
	Draft:       enum.New[StateEnum](1, "草稿"),
	WaitRelease: enum.New[StateEnum](2, "待发布"),
	Published:   enum.New[StateEnum](4, "已发布"),
	Expired:     enum.New[StateEnum](8, "已过期"),
	Revoked:     enum.New[StateEnum](16, "已撤销"),
}

func (e state) New(code int) StateEnum {
	if code == State.Draft.Code() {
		return State.Draft
	}

	if (code & State.WaitRelease.Code()) == State.WaitRelease.Code() {
		return State.WaitRelease
	}

	if (code & State.Published.Code()) == State.Published.Code() {
		return State.Published
	}

	if (code & State.Expired.Code()) == State.Expired.Code() {
		return State.Expired
	}
	if (code & State.Revoked.Code()) == State.Revoked.Code() {
		return State.Revoked
	}

	panic("Announcement.Type.New: error")
}
