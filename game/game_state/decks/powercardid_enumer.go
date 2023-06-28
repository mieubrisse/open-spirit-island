// Code generated by "enumer -type=PowerCardID"; DO NOT EDIT.

package decks

import (
	"fmt"
	"strings"
)

const _PowerCardIDName = "AYearOfPerfectStillnessDrawOfTheFruitfulEarthGuardTheHealingLandRitualsOfDestruction"

var _PowerCardIDIndex = [...]uint8{0, 23, 45, 64, 84}

const _PowerCardIDLowerName = "ayearofperfectstillnessdrawofthefruitfulearthguardthehealinglandritualsofdestruction"

func (i PowerCardID) String() string {
	if i < 0 || i >= PowerCardID(len(_PowerCardIDIndex)-1) {
		return fmt.Sprintf("PowerCardID(%d)", i)
	}
	return _PowerCardIDName[_PowerCardIDIndex[i]:_PowerCardIDIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PowerCardIDNoOp() {
	var x [1]struct{}
	_ = x[AYearOfPerfectStillness-(0)]
	_ = x[DrawOfTheFruitfulEarth-(1)]
	_ = x[GuardTheHealingLand-(2)]
	_ = x[RitualsOfDestruction-(3)]
}

var _PowerCardIDValues = []PowerCardID{AYearOfPerfectStillness, DrawOfTheFruitfulEarth, GuardTheHealingLand, RitualsOfDestruction}

var _PowerCardIDNameToValueMap = map[string]PowerCardID{
	_PowerCardIDName[0:23]:       AYearOfPerfectStillness,
	_PowerCardIDLowerName[0:23]:  AYearOfPerfectStillness,
	_PowerCardIDName[23:45]:      DrawOfTheFruitfulEarth,
	_PowerCardIDLowerName[23:45]: DrawOfTheFruitfulEarth,
	_PowerCardIDName[45:64]:      GuardTheHealingLand,
	_PowerCardIDLowerName[45:64]: GuardTheHealingLand,
	_PowerCardIDName[64:84]:      RitualsOfDestruction,
	_PowerCardIDLowerName[64:84]: RitualsOfDestruction,
}

var _PowerCardIDNames = []string{
	_PowerCardIDName[0:23],
	_PowerCardIDName[23:45],
	_PowerCardIDName[45:64],
	_PowerCardIDName[64:84],
}

// PowerCardIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PowerCardIDString(s string) (PowerCardID, error) {
	if val, ok := _PowerCardIDNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PowerCardIDNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PowerCardID values", s)
}

// PowerCardIDValues returns all values of the enum
func PowerCardIDValues() []PowerCardID {
	return _PowerCardIDValues
}

// PowerCardIDStrings returns a slice of all String values of the enum
func PowerCardIDStrings() []string {
	strs := make([]string, len(_PowerCardIDNames))
	copy(strs, _PowerCardIDNames)
	return strs
}

// IsAPowerCardID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PowerCardID) IsAPowerCardID() bool {
	for _, v := range _PowerCardIDValues {
		if i == v {
			return true
		}
	}
	return false
}
