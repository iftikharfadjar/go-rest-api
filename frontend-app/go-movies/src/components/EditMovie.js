import React, {Component, Fragment} from 'react';
import { Link} from 'react-router-dom';
import './EditMovie.css'
import {Input, TextArea, Select} from './form-components/form-components'

export default class OneGenre extends Component {
	state = {
		movie:{},
		isLoaded:false,
		error: null,
	};

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
		}
		
		this.handleChange = this.handleChange.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
	}

handleSubmit = (evt) => {
	console.log("Form was submitted");
	evt.preventDefault();
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
		
	}


	render() {
		let {movie} = this.state;
		
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
	
					<Input name="title" type="text" title="Title"  value={movie.title} handleChange={this.handleChange}/>
					
		
					
					<Input name="release_date" type="text" title="Release Date"  value={movie.release_date} handleChange={this.handleChange} placeholder="yyyy-mm-dd"/>
					

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