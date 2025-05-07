import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskService } from '../services/task.service';
import { Task } from '../models/task.model';
import { AsyncPipe, NgIf, NgFor } from '@angular/common';
import { TaskEditComponent } from './task-edit.component';
@Component({
  selector: 'app-task-list',
  standalone: true,
  imports: [CommonModule, TaskEditComponent, AsyncPipe, NgIf, NgFor],
  templateUrl: './task-list.component.html',
})
export class TaskListComponent {
  private taskService = inject(TaskService);
  tasks: Task[] = [];

  constructor() {
    this.loadTasks();
  }

  loadTasks() {
    this.taskService.getTasks().subscribe((data) => {
      this.tasks = data;
    });
  }

  selectedTask: Task | null = null;
  showEditModal = false;

  openEditModal(task: Task) {
    this.selectedTask = { ...task }; // crear copia para edición
    this.showEditModal = true;
  }

  closeEditModal() {
    this.showEditModal = false;
    this.selectedTask = null;
  }
  
  deleteTask(id: string) {
    this.taskService.deleteTask(id).subscribe(() => {
      this.loadTasks();
    });
  }
}
