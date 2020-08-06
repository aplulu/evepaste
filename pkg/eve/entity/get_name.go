package entity

func (s *Constellation) GetName(lang string) string {
	for _, lt := range s.Names {
		if *lt.Lang == lang {
			return *lt.Text
		}
	}
	return ""
}

func (s *Region) GetName(lang string) string {
	for _, lt := range s.Names {
		if *lt.Lang == lang {
			return *lt.Text
		}
	}
	return ""
}

func (s *SolarSystem) GetName(lang string) string {
	for _, lt := range s.Names {
		if *lt.Lang == lang {
			return *lt.Text
		}
	}
	return ""
}

func (s *InvCategory) GetName(lang string) string {
	for _, lt := range s.Names {
		if *lt.Lang == lang {
			return *lt.Text
		}
	}
	return ""
}

func (s *InvGroup) GetName(lang string) string {
	for _, lt := range s.Names {
		if *lt.Lang == lang {
			return *lt.Text
		}
	}
	return ""
}

func (s *InvType) GetName(lang string) string {
	for _, lt := range s.Names {
		if *lt.Lang == lang {
			return *lt.Text
		}
	}
	return ""
}

