package ini

import "os"
import "bufio"
import "strings"
import "strconv"
import "github.com/kunos/kml"

type INIReader struct {
	code []string
}

func Open(filename string) (*INIReader, error) {
	fin, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer fin.Close()
	r := bufio.NewReader(fin)

	res := INIReader{}
	for true {
		l, _, err := r.ReadLine()

		if err == nil {
			res.code = append(res.code, string(l))
		} else {
			break
		}
	}
	return &res, nil
}

func (ini *INIReader) GetCode() []string {
	return ini.code
}

func (ini *INIReader) GetSection(name string) []string {
	var start, end int = -1, len(ini.code)

	section := "[" + name + "]"

	for line, ss := range ini.code {

		if strings.Contains(ss, section) {

			start = line
		} else {
			if start != -1 && len(ss) != 0 && ss[0] == '[' {
				end = line
				break
			}
		}
	}

	res := make([]string, 0)
	if start != -1 {
		for i := start + 1; i < end; i++ {
			res = append(res, ini.code[i])
		}
	}

	return res
}

func (ini *INIReader) GetString(section string, key string) string {
	sec := ini.GetSection(section)

	for _, line := range sec {
		if strings.HasPrefix(line, key) {

			splits := strings.Split(line, "=")

			if len(splits) >= 2 {
				comment_splits := strings.Split(splits[1], "//")
				return comment_splits[0]
			}
		}
	}

	return ""
}

func (ini *INIReader) GetBool(section, key string) bool {
	return ini.GetInt(section, key) != 0
}

func (ini *INIReader) GetInt(section string, key string) int {
	ss := strings.TrimSpace(ini.GetString(section, key))

	i, _ := strconv.Atoi(ss)

	return i
}

func (ini *INIReader) GetFloat(section string, key string) float32 {
	ss := strings.TrimSpace(ini.GetString(section, key))

	i, _ := strconv.ParseFloat(ss, 32)

	return float32(i)
}

func (ini *INIReader) GetVec2f(section string, key string) kml.Vec2f {
	str := ini.GetString(section, key)

	splits := strings.Split(str, ",")

	if len(splits) == 2 {

		x, _ := strconv.ParseFloat(strings.TrimSpace(splits[0]), 32)
		y, _ := strconv.ParseFloat(strings.TrimSpace(splits[1]), 32)

		return kml.Vec2f{X: float32(x), Y: float32(y)}
	}
	return kml.Vec2f{}
}

func (ini *INIReader) GetVec3f(section string, key string) kml.Vec3f {
	str := ini.GetString(section, key)

	splits := strings.Split(str, ",")

	if len(splits) == 3 {

		x, _ := strconv.ParseFloat(strings.TrimSpace(splits[0]), 32)
		y, _ := strconv.ParseFloat(strings.TrimSpace(splits[1]), 32)
		z, _ := strconv.ParseFloat(strings.TrimSpace(splits[2]), 32)

		return kml.Vec3f{X: float32(x), Y: float32(y), Z: float32(z)}
	}
	return kml.Vec3f{}
}

func (ini *INIReader) HasSection(section string) bool {
	return len(ini.GetSection(section)) > 0
}

func (ini *INIReader) FuncNoWay32() bool {
	return false
}
