import { NgModule } from '@angular/core';


import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatMenuModule} from '@angular/material/menu';
import {MatGridListModule} from '@angular/material/grid-list';
@NgModule({

imports: [

MatButtonModule,
MatToolbarModule,
MatIconModule,
MatCardModule,
MatInputModule,
MatMenuModule,
MatGridListModule


],

exports: [

MatButtonModule,
MatToolbarModule,
MatIconModule,
MatCardModule,
MatInputModule,
MatMenuModule,
MatGridListModule


]

})

export class MaterialModule {}