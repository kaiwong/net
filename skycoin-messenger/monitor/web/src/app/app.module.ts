import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatGridListModule,
  MatListModule,
  MatIconModule,
  MatTableModule,
  MatTooltipModule,
  MatChipsModule,
  MatSnackBarModule,
  MatCardModule,
  MatButtonModule,
  MatDialogModule,
  MatProgressBarModule,
  MatTabsModule,
  MatFormFieldModule,
  MatInputModule,
  MatProgressSpinnerModule,
  MatMenuModule,
  MatPaginatorModule
} from '@angular/material';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { ApiService, UserService } from './service';
import { TimeAgoPipe, ByteToPipe, EllipsisPipe, IterablePipe, SafePipe } from './pipe';
import { LabelDirective, ShortcutInputDirective, DebugDirective } from './directives';
import { DashboardComponent, SubStatusComponent } from './page';
import {
  UpdateCardComponent,
  AlertComponent,
  LoadingComponent,
  TerminalComponent,
} from './components';
import { AppRoutingModule } from './route/app-routing.module';

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    TimeAgoPipe,
    ByteToPipe,
    EllipsisPipe,
    IterablePipe,
    SafePipe,

    LabelDirective,
    ShortcutInputDirective,
    DebugDirective,

    SubStatusComponent,
    UpdateCardComponent,
    AlertComponent,
    LoadingComponent,
    TerminalComponent,
  ],
  entryComponents: [
    UpdateCardComponent,
    AlertComponent,
    LoadingComponent,
    TerminalComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    MatGridListModule,
    MatListModule,
    MatIconModule,
    MatTableModule,
    MatTooltipModule,
    MatChipsModule,
    MatSnackBarModule,
    MatCardModule,
    MatButtonModule,
    MatDialogModule,
    MatProgressBarModule,
    MatTabsModule,
    MatFormFieldModule,
    MatInputModule,
    MatProgressSpinnerModule,
    MatMenuModule,
    MatPaginatorModule
  ],
  providers: [ApiService, UserService],
  bootstrap: [AppComponent]
})
export class AppModule { }
