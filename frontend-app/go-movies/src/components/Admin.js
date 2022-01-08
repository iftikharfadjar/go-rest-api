import React, {Component, Fragment} from 'react';
import { Link} from 'react-router-dom';

export default class Admin extends Component {
	state = {
		movies:[],
		isLoaded:false,
		error: null
	};
	
	componentDidMount(){
		//fetch("https://backendgo.run-us-west2.goorm.io/v1/movies")
		fetch("http://localhost:700/v1/movies")
		// .then((response) => response.json())  
		.then((response) => {
			console.log("Status code is ", response.status)
			if (response.status !== "200"){
				let err  = Error;
				err.message =  "Invalid response code : " + response.status;
				this.setState({error:err})
			}
			return response.json();
		})
		.then((json) => {
			this.setState({
				movies:json.movies,
				isLoaded:true,
			},
		 (error) => {
				this.setState({
					isLoaded: true,
					error
				})
			}	
		 )
			
		})
	}
	
	render() {
		const {movies,isLoaded, error} = this.state;
		
		if (error) {
			return <div>Error : {error.message}</div>	
		}
		else if(!isLoaded){
			return <p>Loading..</p>
		}else{
			return(
				<Fragment>
					<h2>Choose a Movie</h2>
					
					<div class="list-group">
						{movies.map((m) => (
								<Link key={m.id} to={`admin/movie/${m.id}`}  className="list-group-item list-group-item-action">
									{m.title}
								</Link>
						))}
					</div>
				</Fragment>
			)
		}
	}
	
	
	
}