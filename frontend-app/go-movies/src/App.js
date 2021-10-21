import React from 'react';
import {BrowserRouter as Router, Switch, Route, Link
				// useParams, 
				// useRouteMatch
			 } from 'react-router-dom';
import Movies from './components/Movies'
import OneMovie from './components/OneMovie'
import Genres from './components/Genres'
import OneGenre from './components/OneGenre'
import EditMovie from './components/EditMovie'

function App() {
  return (
		<Router>
		 <div className="container">

				<div className="row">
					<h1 className="mt-3">Go Watch Movie !!!</h1>
					<hr className="mb-3"></hr>
				</div>

				<div className="row">

					<div className="col-md-2">
						<nav>
							<ul className="list-group">
								<li className="list-group-item">
									<Link to="/">Home</Link>
								</li>
								
								<li className="list-group-item">
									<Link to="/movies">Movie</Link>
								</li>
								
								<li className="list-group-item">
									<Link to="/genres">Genres</Link>
								</li>
								
								<li className="list-group-item">
									<Link to="/admin/movie/0">Add Movie</Link>
								</li>
								
								<li className="list-group-item">
									<Link to="/admin">Manage Catalogue</Link>
								</li>
							</ul>
						</nav>
					</div>
					
					
					<div className="col-md-10">
						<Switch>
							<Route path="/movie/:id" component={OneMovie} />
								
							<Route path="/movies">
								<Movies/>
							</Route>
						
							<Route exact path="/genres">
								<Genres/>
							</Route>
							
							<Route path="/genre/:id" component={OneGenre}/>
							
							<Route path="/admin/movie/:id" component={EditMovie} />
							
							<Route path="/admin">
								<Admin/>
							</Route>
							
							<Route path="/">
								<Home/>
							</Route>
						</Switch>
					</div>
					
					
				</div>

			</div>
		</Router>
  );
}

const Home = () => {
	return <h2>Home</h2>
}


const Admin = () => {
	return <h2>Manage Catalogue</h2>
}

export default App;

// const CategoryPage = () => {
// 	let {path,url} = useRouteMatch();
	
// 	return (
// 		<div>
// 			<h2>Category Page</h2>
// 			<ul>
// 				<li><Link to={`${path}/comedy`}>{`Comedy ${url}`}</Link></li>
// 				<li><Link to={`${path}/drama`}>Drama</Link></li>
// 			</ul>
// 		</div>
// 		)
// }

// const Movie = () => {
// 	let { id } = useParams();
// 	return <OneMovie id={}/>
// }


