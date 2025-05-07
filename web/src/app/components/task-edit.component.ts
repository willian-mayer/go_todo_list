import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Task } from '../models/task.model';
import { TaskService } from '../services/task.service';

@Component({
  selector: 'app-task-edit',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './task-edit.component.html',
})
export class TaskEditComponent {
  @Input() task!: Task;
  @Output() close = new EventEmitter<void>();
  @Output() taskUpdated = new EventEmitter<void>();

  constructor(private taskService: TaskService) {}

  updateTask() {
    if (!this.task.title.trim()) return;

    this.taskService.updateTask(this.task.id, this.task).subscribe(() => {
      this.taskUpdated.emit();
      this.close.emit();
    });
  }
}
