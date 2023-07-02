package service

//func (q *QuestionService) UpdateStatus(questionId int, status model.Status) error {
//	chapter := &model.Chapter{}
//	err := db.Mysql.Model(&model.Chapter{}).Where("question_id = ?", questionId).First(&chapter).Error
//	if err != nil {
//		return err
//	}
//
//	chapter.Status = status
//	err = db.Mysql.Model(&model.Chapter{}).Where("id = ?", chapter.ID).Update("status", status).Error
//	if err != nil {
//		return err
//	}
//
//	var count int64
//	err = db.Mysql.Model(&model.Chapter{}).Where("status in (?) and course_id = ?",
//		[]model.Status{
//			model.UnfinishedChapter,
//			model.FailChapter,
//		}, chapter.CourseID).Count(&count).Error
//
//	if err != nil {
//		return err
//	}
//
//	if count == 0 {
//		err = db.Mysql.Model(&model.Course{}).Where("id = ?", chapter.CourseID).Update("status", model.FinishedCourse).Error
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
