package permission

type Permission int64

const (
	Banned        Permission = -999
	Stranger      Permission = -777
	User          Permission = -1
	AdminRDOnly   Permission = 0
	AdminRestrict Permission = 555
	AdminNormal   Permission = 777
	AdminGod      Permission = 999
)

func (this *Permission) String() string {
	switch *this {
	case Banned:
		return "Banned"
	case Stranger:
		return "Stranger"
	case User:
		return "User"
	case AdminRDOnly:
		return "AdminReadOnly"
	case AdminRestrict:
		return "AdminRestrict"
	case AdminNormal:
		return "NormalAdmin"
	case AdminGod:
		return "BossAdmin"
	}
	return `Unknown`
}
