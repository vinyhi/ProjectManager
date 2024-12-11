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

const fetchProjects = async (): Promise<Project[]> => {
  const response = await axios.get(`${API_BASE_URL}/projects`);
  return response.data;
};

const fetchTasks = async (projectId: number): Promise<Task[]> => {
  const response = await axios.get(`${API_BASE_URL}/projects/${projectId}/tasks`);
  return response.data;
};

const createProject = async (project: Project): Promise<Project> => {
  const response = await axios.post(`${API_BASE_URL}/projects`, project);
  return response.data;
};

const updateProject = async (project: Project): Promise<Project> => {
  const response = await axios.put(`${API_BASE_URL}/projects/${project.id}`, project);
  return response.data;
};

const deleteProject = async (id: number): Promise<void> => {
  await axios.delete(`${API_BASE_URL}/projects/${id}`);
};

const createTask = async (task: Task): Promise<Task> => {
  const response = await axios.post(`${API_BASE_URL}/tasks`, task);
  return response.data;
};

const updateTask = async (task: Task): Promise<Task> => {
  const response = await axios.put(`${API_BASE_URL}/tasks/${task.id}`, task);
  return response.data;
};

const deleteTask = async (id: number): Promise<void> => {
  await axios.delete(`${API_BASE_URL}/tasks/${id}`);
};

document.addEventListener("DOMContentLoaded", async () => {
  try {
    const projects = await fetchProjects();
    console.log(projects);
  } catch (error) {
    console.error("Failed to fetch projects", error);
  }
});

document.getElementById("create-project").addEventListener("click", async () => {
  const projectName = (document.getElementById("project-name") as HTMLInputElement).value;
  try {
    const newProject = await createProject({ name: projectName });
    console.log(newProject);
  } catch (error) {
    console.error("Failed to create a project", error);
  }
});