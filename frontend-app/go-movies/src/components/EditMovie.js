import React, {Component, Fragment} from 'react';
import { Link} from 'react-router-dom';
import './EditMovie.css'
import {Input, TextArea, Select} from './form-components/form-components'

export default class OneGenre extends Component {

	constructor(props){
		super(props);
		this.state ={
			movie : {
				id:0,
				title : "",
				release_date : "",
				runtime : "",
				mpaa_rating :"",
				rating : "",
				description : "",
			},
			mpaaOption: [
				{id:"G", value:"G"},
				{id:"PG", value:"PG"},
				{id:"PG13", value:"PG13"},
				{id:"R", value:"R"},
				{id:"NC17", value:"NC17"},
				
			],
			isLoaded : false,
			error: null,
			errors : []
		}
		
		this.handleChange = this.handleChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}

	hasError(key){
		return this.state.errors.indexOf(key) !== -1;
	}
	
handleSubmit = (evt) => {
	console.log("Form was submitted");
	evt.preventDefault();
	
	//client side validation
	let errors = [];
	if(this.state.movie.title === ""){
		errors.push("title")
	}
	
	this.setState({errors : errors});
	
	if (errors.length > 0){
		return false;
	}
	
	const data = new FormData(evt.target);
	const payload = Object.fromEntries(data.entries());
	console.log(payload);
	
	const requestOptions = {
		method : 'POST',
		body : JSON.stringify(payload)
	}
	
	fetch('https://backendgo.run-us-west2.goorm.io/v1/admin/editmovie', requestOptions)
	.then(response => response.json())
	.then(data => {
		console.log(data);
	})
}

handleChange = (evt) => {
	let value = evt.target.value;
	let name = evt.target.name;
	this.setState((prevState) => ({
		movie : {
			...prevState.movie,
			[name] : value,
		}
	}))
}
	
	componentDidMount(){
		const id = this.props.match.params.id;
		if (id > 0) {
			fetch("https://backendgo.run-us-west2.goorm.io/v1/movie/" + id)
			.then((response) => {
				
				if (response.state !== "200"){
					let err = Error;
					err.Message = "Invalid response code: " + response.status;
					this.setState({error: err});	
				}
				return response.json();
			})
			.then((json) => {
				const releaseDate = new Date(json.movie.release_date);
				
				this.setState(
					{
						movie: {
							id : id,
							title : json.movie.title,
							release_date : releaseDate.toISOString().split("T")[0],
							runtime : json.movie.runtime,
							mpaa_rating : json.movie.mpaa_rating,
							rating : json.movie.rating,
							description : json.movie.description,
						},
						isLoaded: true,
					},
					(error) => {
						this.setState({
							isLoaded: true,
							error,
						})
					}
				)
			})
		} else {
			this.setState({isLoaded: true})
		}
	}


	render() {
		let {movie, isLoaded, error} = this.state;
		
		if (error) {
			return <div>Error: {error.message}</div>
		}else if(!isLoaded) {
			return <p>Loading...</p>
		}
		else{
			return(
				<Fragment>
					<h2>Add/Edit Movie</h2>	
					<hr/>
					<form onSubmit={this.handleSubmit}>
						<input 
							type="hidden"
							name="id"
							id="id"
							value={movie.id}
							onChange={this.handleChange}
							/>

						<Input 
							className={this.hasError("title") ? "is-invalid" : ""} 
							errorDiv={this.hasError("title") ? "text-danger" : "d-none"} 
							errorMsg={"Please enter a title"}
							name="title" type="text" title="Title"  value={movie.title} handleChange={this.handleChange}/>



						<Input name="release_date" type="date" title="Release Date"  value={movie.release_date} handleChange={this.handleChange}/>


						<Input name="runtime" type="text" title="Runtime"  value={movie.runtime} handleChange={this.handleChange}/>



						<Select name="mpaa_rating" title="MPAA Rating" value={movie.mpaa_rating} handleChange={this.handleChange} option={this.state.mpaaOption} ></Select>

						<Input name="rating" type="text" title="Rating"  value={movie.rating} handleChange={this.handleChange}/>


						<TextArea name="description" handleChange={this.handleChange} value={movie.description} title="Description" />

						<hr/>

						<button className="btn btn-primary">Save</button>
					</form>

					<div className="mt-3">
						<pre>{JSON.stringify(this.state, null, 3)}</pre>
					</div>

				</Fragment>
				)
		}
	}

}