import {Component, Injectable, OnInit} from '@angular/core';
import { UploadFilesService } from 'src/app/services/upload.service';
import {HttpEvent, HttpEventType, HttpProgressEvent, HttpResponse} from '@angular/common/http';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-upload-files',
  templateUrl: './upload.component.html',
})
@Injectable({
  providedIn: 'root'
})
export class UploadFilesComponent implements OnInit {
  selectedFiles?: FileList;
  progressInfos: any[] = [];
  message: string[] = [];
  fileInfos?: Observable<any>;
  constructor(private uploadService: UploadFilesService) { }

  ngOnInit() {
    this.fileInfos = this.uploadService.getFiles();
  }

  selectFiles(event:any): void {
    this.message = [];
    this.progressInfos = [];
    this.selectedFiles = event.target.files;
    console.log(this.selectedFiles)
    console.log(this)
  }

  getSelectedFiles() : FileList {
    return this.selectedFiles!;
  }

  uploadFiles(): void {
    this.message = [];
    if (this.selectedFiles) {
      for (let i = 0; i < this.selectedFiles.length; i++) {
        this.upload(i, this.selectedFiles[i]);
      }
    }
  }

  upload(idx: number, file: File): void {
    this.progressInfos[idx] = { value: 0, fileName: file.name };
    if (file) {
      this.uploadService.upload(file).subscribe(
        (event: any) => {
          if (event.type === HttpEventType.UploadProgress) {
            this.progressInfos[idx].value = Math.round(100 * event.loaded / event.total);
          } else if (event instanceof HttpResponse) {
            const msg = 'Uploaded the file successfully: ' + file.name;
            this.message.push(msg);
            this.fileInfos = this.uploadService.getFiles();
          }
        },
        (err: any) => {
          this.progressInfos[idx].value = 0;
          const msg = 'Could not upload the file: ' + file.name;
          this.message.push(msg);
          this.fileInfos = this.uploadService.getFiles();
        });
    }
  }
}
