import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskListComponent } from './components/task-list.component';
import { TaskCreateComponent } from './components/task-create.component';
@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, TaskListComponent, TaskCreateComponent],
  templateUrl: './app.component.html'
})

export class AppComponent {
  showModal = false;

  openModal() {
    this.showModal = true;
  }

  closeModal() {
    this.showModal = false;
  }

  reloadTasks() {
    // Aquí puedes comunicarte con TaskListComponent si necesitas recargar
    // o usar un servicio compartido para emitir eventos
    location.reload(); // simple para ahora
  }


}