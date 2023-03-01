# Documentation

From request api gateway, generate api-key token and register to specific usage \
plan in api gateway, if the email exist in dynamodb, return error, else continue.\
the api key token is sent to user email by ses service.

## usage

write file env.json with all env variables, and run: \
``` terraform init ``` \
``` terraform apply -var-file=env.json ``` \

need template email uploaded to ses, upload with script
``` bash upload_template.sh ```
need file template.json with data:

``` json
{
    "Template":{
        "TemplateName": "string",
        "SubjectPart": "string",
        "TextPart": "string",
        "HtmlPart": "string"
    }
}
```

### env variables.

- "DynamoDBTable": "string",
- "USAGE_PLAN_ID": "string",
- "EmailToNotifyErrorsApi": "string",
- "EMAIL_SOURCE": "string",
- "EMAIL_SUBJECT": "string",
- "EMAIL_TEST": "string",
- "LAMBDA_FUNCTION_NAME": "string",
- "ROLE_NAME": "string",
- "TEMPLATE_EMAIL_SES_NAME": "string"