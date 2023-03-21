import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { StudentListComponent } from './student-list.component';
import { StudentQueryComponent } from './student-query.component';

const routes: Routes = [
  { path: 'list', component: StudentListComponent },
  { path: 'query', component: StudentQueryComponent }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class StudentRoutingModule {}
