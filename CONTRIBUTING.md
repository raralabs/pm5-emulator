**Contributing**
-------------
When contributing to this repository, please first discuss the change you wish to make via issue, email, or any other method with the owners of this repository before making a change.

Please follow [Code of Conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct/) in all your interactions with the project.

**Contributor Workflow**
To contribute, the workflow is as follows:
- Fork repository
- Create topic branch
- Commit your change
- Push changes to your fork
- Create pull request

**Pull Request Process**
1. Ensure your fork is upto date. 
2. Update README.md with details of changes to interface, this includes new environment variables, exposed ports, useful file locations,etc.
3. Your pull request will be merged once one of the reviewer reviews the changes. 


**Some tips for commits**
1. Write quality commit messages.
2. Do not add to many redundant commits. If so, squash your commits using:
```cassandraql
git checkout your-branch
git rebase -i HEAD~n 
# n is the number of commits in pull request 
git push #push to github 
``` 

