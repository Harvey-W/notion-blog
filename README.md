# NotionCMS-for-Hugo 使用Notion写文章同步更新Hugo博客

`NotionCMS-for-Hugo` allows you to use Notion as a CMS for pages built with hugo. You can use it as a cli or even automate your blog repo to update itself with the Github Action.  
<u>NOTE: This is based on [notion-blog](https://github.com/xzebra/notion-blog), which is no longer maintained by the developer.</u>  

Demo: [![](https://img.shields.io/badge/Harvey's%20Blog-@HarveyW-blue)](https://github.com/Harvey-W/harvey-w.github.io)

## Changelog

- Fixed a warning caused by a deprecated GitHub Actions command:  
  ```JavaScript
  Warning: The `set-output` command is deprecated and will be disabled soon. Please upgrade to using Environment Files. For more information see: https://github.blog/changelog/2022-10-11-github-actions-deprecating-save-state-and-set-output-commands/  
- Replaced title-based identification with Notion ID for post synchronization.
Now all add/delete/modify operations are tracked using the Notion ID instead of the post title.
- Deletions in the Notion database are now fully synchronized with corresponding posts and attached images, thanks to the Notion ID–based identifier.
- Deprecated the "Author" property, as it is not included by default in new Notion pages and may cause critical errors when missing.
- Updated all tools in the workflow to their latest versions and added more descriptive Git commit messages.   
  https://github.com/Harvey-W/harvey-w.github.io/blob/main

## Usage

- YOUR Notion Integration Secret (GitHub repository -> Settings -> Secrets and variables -> Actions -> Repository secrets -> New)
- YOUR Notion Database shared ID (/.github/workflows/notion-sync.yml -> "databaseID":)
- Hugo powered blog.
