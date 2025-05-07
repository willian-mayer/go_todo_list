import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskService } from '../services/task.service';
import { Task } from '../models/task.model';
import { AsyncPipe, NgIf, NgFor } from '@angular/common';

@Component({
  selector: 'app-task-list',
  standalone: true,
  imports: [CommonModule, AsyncPipe, NgIf, NgFor],
  templateUrl: './task-list.component.html' 
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
}
