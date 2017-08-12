package room

import (
	//"fmt"
	"server/conf"
	"server/db/mysql"
	"server/msg"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
	字段
     'room_id' => $room_id,
     'role_id' => $role_id,
     'stage_id' => $stage_id,
     'stage_level'=>$stage_info->lv,
     'user_num'=>1,
     'limit_level' => $limit_level,
     'is_open' => $is_open,
     'status' => ROOM_FREE_STATUS,
     'fighting_power' =>
             array(
                 $role_id=>$ut['fighting_power']
             ),
     'user_status'=>
             array(
                 $role_id=>USER_ROOM_UNREADY_STATUS
             ),
     'room_users' => (array)$this->user_role_model->role_game_list($role_id),
     'teams' =>
             array(
                 $role_id=>$team_info,
             ),
     'simple_teams'=>
             array(
                 $role_id=>$simple_teams,
             )
*/

type Room struct {
	Room_id int `gorm:"primary_key"`

	Role_id        int
	Stage_id       int
	Stage_level    int
	User_num       int
	Limit_level    int
	Is_open        int
	Status         int
	Fighting_power string //json for map[int]map[string]string
	User_status    string //json for map[int]map[string]string
	Room_users     string //json for map[int]map[string]string
	Teams          string //json for map[int]map[string]string
	Simple_teams   string //json for map[int]map[string]string
}

func (d Room) TableName() string {
	return "soosoogoo_rooms"
}

/**
user info 查询结构体
*/
type RoomModel struct {

	//继承mysql.MysqlDriver
	mysql.MysqlDriver
}

//#TODO::构造函数怎么写

//查询房间信息
func (u *RoomModel) GreateRoom(m *msg.CreateRoom) Room {

	var uDB = u.Connent(conf.SOCKETDATABASE)

	room := Room{
		Role_id:        m.Stage_id,
		Stage_id:       m.Stage_id,
		Stage_level:    m.Stage_id,
		User_num:       1,
		Limit_level:    1,
		Is_open:        m.Is_open,
		Status:         m.Stage_id,
		Fighting_power: "",
		User_status:    "",
		Room_users:     "",
		Teams:          "",
		Simple_teams:   "",
	}

	uDB.Create(&room)

	return room
}
