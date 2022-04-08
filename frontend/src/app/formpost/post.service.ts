import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class PostService {

  constructor(private http: HttpClient) {
  }

  getPosts(): Observable<Response> {
    const result = this.http.get<Response>(environment.gateway + '/post');
    console.log("GetPosts", result)
    // result.subscribe(data => {
    //   console.log(data)
    // })
    return result
  }

  fetchCategories() : Observable<any> {
    return this.http.get<any>(environment.gateway + '/categories')
  }

  fetchSubcategories(categoryId : number) : Observable<any> {
    return this.http.get<any>(environment.gateway + '/subcategories/' + categoryId)
  }

  fetchLocations() : Observable<any> {
    return this.http.get<any>(environment.gateway + '/locations')
  }

  createPost(jobPost: JobPost, files?: FileList) : Observable<any> {
    const formData = new FormData();
    formData.append('data', JSON.stringify(jobPost))
    if (files != null) {
      for (let i = 0; i < files.length; i++) {
        formData.append('files', files![i], files![i].name);
      }
    }
    console.log(formData.getAll('data'))
    return this.http.post<any>(environment.gateway + '/post/category/' + jobPost.Post.Categoryid, formData)
  }
}


export interface Categories {
  CategoryId:number;
  Name:string;
}

export interface Subcategories {
  SubcategoryId:number;
  CategoryId:number;
  Name:string;
}

export interface Location {
  Name:string;
}

export interface Post {
  Sellerid: number;
  Categoryid: number;
  Subcategoryid: number;
  Title: string;
  Description: string;
  CreatedAt?:string;
  Image?:FileList;
}

export interface Job {
  Salary:number;
  Pay:string;
  Type:string;
  Location:string;
  Place:string;
}

export interface JobPost {
  Post:Post;
  Job:Job;
}

export interface Response {
   data: Array<Post>
   message: string
   status : number
}