# NotionCMS-for-Hugo 使用Notion写文章同步更新Hugo博客

`NotionCMS-for-Hugo` allows you to use Notion as a CMS for pages built with hugo. You can use it as a cli or even automate your blog repo to update itself with the Github Action.  
<u>NOTE: This is based on [notion-blog](https://github.com/xzebra/notion-blog), which is no longer maintained by the developer.</u>  

Demo: [![](https://img.shields.io/badge/Harvey's%20Blog-@HarveyW-blue)](https://github.com/Harvey-W/harvey-w.github.io)

## Changelog

- Fix a warning caused by a GitHub Action deprecated command:  
  ```JavaScript
  Warning: The `set-output` command is deprecated and will be disabled soon. Please upgrade to using Environment Files. For more information see: https://github.blog/changelog/2022-10-11-github-actions-deprecating-save-state-and-set-output-commands/  
- Use **Notion ID** to identify an unique post and apply on "add/delete/modify". (Instead of title name in the original version).
- Delete pages in the Notion Database can NOW be synchronized equally with posts and attached images due to the Notion ID identifier.
- Deprecate "Author" property because this is not a default setting in a new Notion page, and can cause a severe bug when it is missing.
- All tools are update to latest in the workflow and add an elaborate git commit message.  
  https://github.com/Harvey-W/harvey-w.github.io/blob/main

## Usage

- YOUR Notion Integration Secret (GitHub repository -> Settings -> Secrets and variables -> Actions -> Repository secrets -> New)
- YOUR Notion Database shared ID (/.github/workflows/notion-sync.yml -> "databaseID":)
- Hugo powered blog.
