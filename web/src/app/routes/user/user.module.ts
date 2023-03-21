import { NgModule, Type } from '@angular/core';
import { SharedModule } from '@shared';

import { UserListComponent } from './user-list.component';
import { UserRoutingModule } from './user-routing.module';

const COMPONENTS: Array<Type<void>> = [UserListComponent];

@NgModule({
  imports: [SharedModule, UserRoutingModule],
  declarations: COMPONENTS
})
export class UserModule {}
