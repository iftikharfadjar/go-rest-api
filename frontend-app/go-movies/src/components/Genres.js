import React, {Component, Fragment} from 'react';
import { Link} from 'react-router-dom';

export default class Genres extends Component {
	state = {
		genres:[],
		isLoaded:false,
		error: null
	};
	
	componentDidMount(){
		//fetch("https://backendgo.run-us-west2.goorm.io/v1/genres")
		fetch("http://localhost:700/v1/genres")
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
				genres:json.genres,
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
		const {genres,isLoaded, error} = this.state;
		
		
		if (error) {
			return <div>Error : {error.message}</div>	
		}
		else if(!isLoaded){
			return <p>Loading..</p>
		}else{
			return(
				<Fragment>
					<h2>Genres</h2>
					
					<div class="list-group">
						{genres.map((g) => (
								<Link 
									key={g.id}
									className="list-group-item list-group-item-action"
									to={{
										pathname: `/genre/${g.id}`,
										genreName: g.genre_name,
									}}>{g.genre_name}</Link>
						))}
						
					</div>
				</Fragment>
			)
		}
	}
	
	
	
}