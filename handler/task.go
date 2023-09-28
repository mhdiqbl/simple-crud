package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ottoDigital/dto"
	"ottoDigital/model"
	"ottoDigital/utils"
)

func (h *handler) CreateTask(c *gin.Context) {
	var taskRequest *dto.TaskRequest
	var task model.Task

	if err := c.ShouldBind(&taskRequest); err != nil {
		c.JSON(400, err)
		return
	}

	task.UserId = taskRequest.UserId
	task.Tittle = taskRequest.Tittle
	task.Description = taskRequest.Description
	task.Status = taskRequest.Status

	taskId, err := h.repository.CreateTask(c, task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil create task",
		Data:    taskId,
	})
}

func (h *handler) GetTasks(c *gin.Context) {
	taskResponse := make([]*dto.TaskResponse, 0)
	tasks, err := h.repository.GetAllTasks(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	for _, task := range tasks {
		taskResponse = append(taskResponse, &dto.TaskResponse{
			ID:          task.ID,
			UserId:      task.UserId,
			Tittle:      task.Tittle,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdateAt:    task.UpdateAt,
		})
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil get data tasks",
		Data:    &taskResponse,
	})
}

func (h *handler) GetTaskByID(c *gin.Context) {
	var taskResponse *dto.TaskResponse
	taskID, err := utils.ExtractorRequestParamID(&c.Params)

	task, err := h.repository.GetTaskByID(c, int(taskID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
		return
	}

	if task.ID == 0 {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Candidate not found",
			Data:    nil,
		})
		return
	}

	taskResponse = &dto.TaskResponse{
		ID:          task.ID,
		UserId:      task.UserId,
		Tittle:      task.Tittle,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdateAt:    task.UpdateAt,
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil get data task",
		Data:    &taskResponse,
	})
	return
}

func (h *handler) UpdateTask(c *gin.Context) {
	var task model.Task
	var taskRequest *dto.TaskRequest

	taskId, err := utils.ExtractorRequestParamID(&c.Params)

	if err != nil {
		c.JSON(400, err)
		return
	}

	if err := c.ShouldBind(&taskRequest); err != nil {
		c.JSON(400, err)
		return
	}

	task.UserId = taskRequest.UserId
	task.Tittle = taskRequest.Tittle
	task.Description = taskRequest.Description
	task.Status = taskRequest.Status

	var taskFound bool
	//fmt.Println(user)
	taskFound, err = h.repository.UpdateTask(c, int(taskId), task)

	if err != nil {
		c.JSON(400, err)
		return
	}

	if !taskFound {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Task not found",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success update task",
		Data:    nil,
	})
	return
}

func (h *handler) DeleteTask(c *gin.Context) {
	taskId, err := utils.ExtractorRequestParamID(&c.Params)

	if err != nil {
		c.JSON(400, err)
		return
	}

	var taskFound bool
	taskFound, err = h.repository.DeleteTask(c, int(taskId))

	if !taskFound {
		c.JSON(http.StatusNotFound, dto.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Task not found",
			Data:    nil,
		})
		return
	}

	if err != nil {
		c.JSON(400, err)
		return
	}

	c.JSON(http.StatusOK, dto.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success delete task",
		Data:    nil,
	})
	return
}
