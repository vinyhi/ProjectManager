import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_BACKEND_URL;

interface Project {
  id?: number;
  name: string;
}

interface Task {
  id?: number;
  name: string;
  projectId: number;
}

const makeApiCall = async <T>(method: "get" | "post" | "put" | "delete", url: string, data?: Project | Task): Promise<T> => {
  const config = {
    method,
    url: `${API_BASE_URL}${url}`,
    data,
  };
  
  const response = await axios(config);
  return response.data;
};

const fetchProjects = async (): Promise<Project[]> => {
  return makeApiCall<Project[]>('get', '/projects');
};

const fetchTasks = async (projectId: number): Promise<Task[]> => {
  return makeApiCall<Task[]>('get', `/projects/${projectId}/tasks`);
};

const createProject = async (project: Project): Promise<Project> => {
  return makeApiCall<Project>('post', '/projects', project);
};

const updateProject = async (project: Project): Promise<Project> => {
  return makeApiCall<Project>('put', `/projects/${project.id}`, project);
};

const deleteProject = async (id: number): Promise<void> => {
  await makeApiCall<void>('delete', `/projects/${id}`);
};

const createTask = async (task: Task): Promise<Task> => {
  return makeApiCall<Task>('post', '/tasks', task);
};

const updateTask = async (task: Task): Promise<Task> => {
  return makeApiCall<Task>('put', `/tasks/${task.id}`, task);
};

const deleteTask = async (id: number): Promise<void> => {
  await makeApiCall<void>('delete', `/tasks/${id}`);
};

document.addEventListener("DOMContentLoaded", async () => {
  try {
    const projects = await fetchProjects();
    console.log(projects);
  } catch (error) {
    console.error("Failed to fetch projects", error);
  }
});

document.getElementById("create-project")?.addEventListener("click", async () => {
  const projectNameElement = document.getElementById("project-name") as HTMLInputElement;
  if (!projectNameElement) return;
  const projectName = projectNameElement.value;
  try {
    const newProject = await createProject({ name: projectName });
    console.log(newProject);
  } catch (error) {
    console.error("Failed to create a project", error);
  }
});