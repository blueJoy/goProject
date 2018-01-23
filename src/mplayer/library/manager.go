package library

import "errors"

type MusicEntry struct{
	Id string
	Name string
	Artist string
	Source string
	Type string
}

type MusicManager struct {
	musics []MusicEntry
}

//初始化函数
func NewMusicManager() *MusicManager{
	return &MusicManager{make([]MusicEntry,0)}
}

//获取切片数组的长度
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//根据index位置获取数组对应的music对象
func (m *MusicManager) Get(index int) (music *MusicEntry,err error){
	if index < 0 || index > m.Len() {
		return nil,errors.New("Index out of range.")
	}
	return &m.musics[index],nil
}

//根据音乐名称获取对象
func (m *MusicManager) Find(name string)  *MusicEntry {
	if m.Len() == 0 {
		return nil
	}

	for _,m := range  m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}

//添加音乐
func (m *MusicManager) Add(music *MusicEntry){
	//疑问：为什么append的music 也要加*
	m.musics=append(m.musics,*music)
}

//删除音乐
func (m *MusicManager) Remove(index int) *MusicEntry{
	if m.Len() < 0 || index > m.Len() {
		return nil
	}

	removedMusic := &m.musics[index]

	//从数组切片中删除元素
	if index < m.Len() -1 {
		//删除中间元素
		m.musics = append(m.musics[:index-1],m.musics[index + 1:]...)
	}else if index ==0 {
		//只有一个元素
		m.musics = make([]MusicEntry,0)
	}else{
		//删除最后一个元素
		m.musics = m.musics[:index]
	}
	return removedMusic
}

//按照名称删除
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if m.Len() == 0{
		return nil
	}

	for i,v := range m.musics{
		if v.Name == name{
			return m.Remove(i)
		}
	}
	return nil
}

