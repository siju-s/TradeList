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
    return this.http.get<Response>(environment.gateway + '/post')
  }

  getPostsForSubcategory(id : number): Observable<Response> {
    return this.http.get<Response>(environment.gateway + '/post/subcategory/' + id)
  }
  
  getPostsForUserId(id : number): Observable<Response> {
    return this.http.get<Response>(environment.gateway + '/post/user/' + id)
  }

  fetchCategories() : Observable<any> {
    return this.http.get<any>(environment.gateway + '/categories');
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
    return this.http.post<any>(environment.gateway + '/post/category/' + jobPost.Post.CategoryId, formData)
  }

  getPayTypes() : Array<any> {
    return [
      {label: 'hourly', value: 'hourly'},
      {label: 'monthly', value: 'monthly'},
      {label: 'yearly', value: 'yearly'}
      ]
  }

  getJobTypes() : Array<any> {
    return [
      {label: 'fulltime', value: 'fulltime'},
      {label: 'partime', value: 'partime'},
      {label: 'internship', value: 'internship'}
    ]
  }

  getLocationTypes() : Array<any> {
    return [
      {label: 'onsite', value: 'onsite'},
      {label: 'remote', value: 'remote'},
      {label: 'hybrid', value: 'hybrid'}
    ]
  }

  instanceOfJobPost(data: any): data is JobPost {
    return 'Post' in data && 'Job' in data;
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
  ID?:number;
  SellerId: number;
  CategoryId: number;
  SubcategoryId: number;
  Title: string;
  Description: string;
  CreatedAt?:string;
  Image?:Array<Image>;
  PostedBy?:string;
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
   data: Array<any>
   message: string
   status : number
}

export interface Image {
  ID:number;
  PostId:number;
  SellerId: number;
  Url:string;
}
