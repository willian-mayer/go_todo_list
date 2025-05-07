import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TaskService } from '../services/task.service';

@Component({
  selector: 'app-task-create',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './task-create.component.html',
  styles: [`
    .modal {
      position: fixed;
      top: 0; left: 0;
      width: 100vw; height: 100vh;
      background: rgba(0,0,0,0.5);
      display: flex; align-items: center; justify-content: center;
    }
    .modal-content {
      background: white;
      padding: 2rem;
      border-radius: 8px;
      width: 400px;
    }
  `]
})
export class TaskCreateComponent {
  title = '';
  content = '';

  @Output() close = new EventEmitter<void>();
  @Output() taskCreated = new EventEmitter<void>();

  constructor(private taskService: TaskService) {}

  createTask() {
    if (!this.title.trim()) return;

    this.taskService.createTask({
      title: this.title,
      content: this.content,
      is_done: false
    }).subscribe(() => {
      this.taskCreated.emit(); // Notifica al padre para recargar la lista
      this.close.emit();       // Cierra el modal
    });
  }
}
