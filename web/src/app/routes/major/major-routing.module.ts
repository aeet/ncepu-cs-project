import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { MajorListComponent } from './list/list.component';

const routes: Routes = [
  { path: 'list', component: MajorListComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MajorRoutingModule {}
