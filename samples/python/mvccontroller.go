package pythonsamples

type MvcController struct {}

var mvcImport *AppImport

func (mvc * MvcController) TaskController()  string {
return ``+mvcImport.ImportForTaskController()+`

class TaskController:
	pass`
}

func (mvc * MvcController) UserController()  string {
return ``+mvcImport.ImportForUserController()+`

class UserController:
	pass`
}

func (mvc * MvcController) RoleController()  string {
return ``+mvcImport.ImportForAllControllers()+`

class RoleController:

	'''@app.route(f'/empresas', methods=['GET'])
	def listar_empresas():
		empresas = Empresa.obter_todos()
		qtd = len(empresas)
		return render_template('empresa/listar.html', lista_empresas=empresas, num_registos=qtd,
			usuario_logado=Usuario.obter_usuario_logado())


	@app.route(f'/empresas/<id>/detalhes', methods=['GET'])
	def detalhes_empresa(id):
		empresa = Empresa.obter_por_unique_id(id)
		if empresa:
			return render_template('empresa/detalhes.html', empresa=empresa, usuario_logado=Usuario.obter_usuario_logado())
		else:
			return render_template('erro/404.html')


	@app.route(f'/empresas/registar', methods=['GET', 'POST'])
	def registar_empresa():
		if request.method == 'GET': 
			return render_template('empresa/registar.html', usuario_logado=Usuario.obter_usuario_logado())
		
		nome_empresa = request.form['nome_empresa']
		nif = request.form['nif']
		telefone = request.form['telefone']
		email = request.form['email']
		empresa = Empresa(nome_empresa, nif)
		empresa.salvar()
		contacto = ContactoEmpresa(empresa.id_empresa, telefone, email)
		contacto.salvar()
		return redirect(url_for('listar_empresas'))


	@app.route(f'/empresas/pesquisar', methods=['GET', 'POST'])
	def pesquisar_empresas():
		if request.method == 'GET': 
			return render_template('empresa/pesquisar.html',usuario_logado=Usuario.obter_usuario_logado())
		else:
			valor = request.form['valor_pesquisa']
			res = Empresa.buscar(valor)
			num_registos =  len(res)
			return render_template('empresa/resultado-pesquisa.html', param=valor, resultados=res, 
					num_registos=num_registos, usuario_logado=Usuario.obter_usuario_logado())'''`
}

func (mvc * MvcController) AuthController()  string {
return ``+mvcImport.ImportForTaskController()+`

class AuthController:

	@app.route('/', methods=['GET'])
	@app.route('/login', methods=['GET', 'POST'])
	def login():
		if request.method == 'GET':
			return render_template('auth/login.html')
		else:
			user_name = request.form['user_name']
			password = request.form['password']
			if(User.exists(user_name, password)):
				session['user_name'] = user_name
				session['password'] = password
				return redirect(url_for('home'))
			else:
				return redirect(url_for('login'))

	@app.route('/logout', methods=['GET'])
	def logout():
		session.pop('user_name')
		return redirect(url_for('login'))

	@app.route('/home', methods=['GET'])
	def home():
		return render_template('auth/home.html', logged_user=User.get_logged_user())`
}