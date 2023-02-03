interface Repository {
  find: (id: number) => Promise<string>;
  delete: (id: number) => Promise<number>;
}

class RepositoryImpl implements Repository {
  find(id: number): Promise<string> {
    return Promise.resolve("foo");
  }
  delete(id: number): Promise<number> {
    return Promise.resolve(1);
  }
}

class Service {
  constructor(private repository: Repository) { }
  async doSomething(id: number) {
    const result = await this.repository.find(id);
    const deleteRes = await this.repository.delete(id);
  }
}