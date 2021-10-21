const Input = (props) => {
	return(
		<div className="mb-3">
			<label htmlFor={props.name} className="form-label">{props.title}</label>
			<input type={props.type} 
				className={`form-control ${props.className}`}
				id={props.name} 
				name={props.name} 
				value={props.value} 
				onChange={props.handleChange} 
				placeholder={props.placeholder}
				/>
			<div className={props.errorDiv}>{props.errorMsg}</div>
		</div>
	);
}

const TextArea = (props) => {
	return(
			<div className="mb-3">
				<label htmlFor={props.name} className="form-label">{props.title}</label>
				<textarea 
					className="form-control" 
					id={props.name} 
					name={props.name} 
					rows="3" 
					onChange={props.handleChange} 
					value={props.value} />
			</div>
	);
}

const Select = (props) => {
	return(
		<div className="mb-3">
						<label htmlFor={props.name} className="form-label">{props.title}</label>
						<select className="form-select" id={props.name} name={props.name} value={props.value} onChange={props.handleChange} >
							<option className="form-select" value="">{props.placeholder}</option>
							{props.option.map((opt) => {
								return(
									<option 
										className="form-select"
										key={opt.id}
										value={opt.id}
										label={opt.value}
										>
										{opt.value}
									</option>
								)
							})}
						</select>
					</div>
	)
}

export {TextArea, Input, Select};

