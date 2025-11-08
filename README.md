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
- Updated all toolkits in the [workflow](https://github.com/Harvey-W/harvey-w.github.io/blob/main/.github/workflows/notion-sync.yml) to their latest versions and added more descriptive Git commit messages.   
  

## Usage 使用方法

- Apply YOUR [Notion Integration Secret](https://www.notion.so/profile/integrations) and connect to your Notion Page(Database needed). Follow with: GitHub repository -> Settings -> Secrets and variables -> Actions -> Repository secrets -> New  
  *申请一个Notion Integration，并连接到所在Notion页面（要用数据库形式存放文章），然后依照顺序添加此secret*
- Copy YOUR Notion Database shared ID, which is contained in "Shared Page" of Notion page(formed like https://www.notion.so/YOUR_ID?v=xxxx). Follow with: /.github/workflows/notion-sync.yml -> "databaseID":  
 *复制分享ID，通过分享一个Notion页面来找到ID，然后粘贴到workflow里文件对应键值*
- Hugo powered blog.
