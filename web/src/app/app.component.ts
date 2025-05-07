import { Component } from '@angular/core';
import { TaskListComponent } from './components/task-list.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [TaskListComponent],
  templateUrl: './app.component.html'
})
export class AppComponent {}
