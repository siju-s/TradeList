import { NgModule } from '@angular/core';


import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatMenuModule} from '@angular/material/menu';
@NgModule({

imports: [

MatButtonModule,
MatToolbarModule,
MatIconModule,
MatCardModule,
MatInputModule,
MatMenuModule,


],

exports: [

MatButtonModule,
MatToolbarModule,
MatIconModule,
MatCardModule,
MatInputModule,
MatMenuModule,


]

})

export class MaterialModule {}