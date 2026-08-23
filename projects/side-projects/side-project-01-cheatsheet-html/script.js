const commands = [
  {name:'init', alias:'git init', desc:'Create a new local repository', options:['--bare','--quiet','--template=<template>','--separate-git-dir']},
  {name:'clone', alias:'git clone', desc:'Clone a repository into a new directory', options:['--depth <n>','--branch <name>','--recurse-submodules','--single-branch']},
  {name:'status', alias:'git status', desc:'Show the working tree status', options:['--short','--branch','--show-stash','--porcelain']},
  {name:'add', alias:'git add', desc:'Add file contents to the index', options:['<pathspec>','-p / --patch','-A / --all','-n / --dry-run']},
  {name:'commit', alias:'git commit', desc:'Record changes to the repository', options:['-m <msg>','--amend','-a / --all','--no-edit']},
  {name:'push', alias:'git push', desc:'Update remote refs and send objects', options:['<remote> <branch>','--force','--set-upstream','--tags']},
  {name:'pull', alias:'git pull', desc:'Fetch from and integrate with another repository', options:['--rebase','--ff-only','--no-commit','--autostash']},
  {name:'fetch', alias:'git fetch', desc:'Download objects and refs from another repository', options:['--all','--prune','--depth <n>','--tags']},
  {name:'branch', alias:'git branch', desc:'List, create, or delete branches', options:['-d / --delete','-D','-m / --move','--list']},
  {name:'checkout', alias:'git checkout', desc:'Switch branches or restore working tree files', options:['-b <branch>','--detach','--patch','--ours/--theirs']},
  {name:'merge', alias:'git merge', desc:'Join two or more development histories together', options:['--no-ff','--squash','--abort','--ff-only']},
  {name:'rebase', alias:'git rebase', desc:'Reapply commits on top of another base tip', options:['-i / --interactive','--onto <newbase>','--continue','--abort']},
  {name:'log', alias:'git log', desc:'Show commit logs', options:['--oneline','--graph','--stat','-p']},
  {name:'reset', alias:'git reset', desc:'Reset current HEAD to the specified state', options:['--soft','--mixed','--hard','--keep']},
  {name:'stash', alias:'git stash', desc:'Stash the changes in a dirty working directory', options:['save','pop','list','apply']},
  {name:'remote', alias:'git remote', desc:'Manage set of tracked repositories', options:['-v','add <name> <url>','remove <name>','show <name>']},
  {name:'tag', alias:'git tag', desc:'Create, list, delete or verify a tag object', options:['-a <name>','-d <name>','-v <name>','--list']},
  {name:'revert', alias:'git revert', desc:'Revert some existing commits', options:['-n / --no-commit','-m <parent-number>','--no-edit','--edit']}
];

const container = document.getElementById('commands');
const searchInput = document.getElementById('search');

function renderCards(list){
  container.innerHTML = '';
  list.forEach(cmd => container.appendChild(cardFor(cmd)));
}

function cardFor(cmd){
  const el = document.createElement('section'); el.className='card';
  const h = document.createElement('h2');
  const nameSpan = document.createElement('span'); nameSpan.className='cmd'; nameSpan.textContent = cmd.alias || cmd.name;
  h.appendChild(nameSpan);
  const d = document.createElement('div'); d.className='desc'; d.textContent = cmd.desc;
  el.appendChild(h); el.appendChild(d);

  const opts = document.createElement('div'); opts.className='options';
  const top = cmd.options.slice(0,3);
  top.forEach(o => opts.appendChild(optionNode(o)));

  if(cmd.options.length>3){
    const more = document.createElement('div'); more.className='more hidden';
    cmd.options.slice(3).forEach(o => more.appendChild(optionNode(o)));
    const btn = document.createElement('button'); btn.className='more-toggle'; btn.textContent='More ▾';
    btn.setAttribute('aria-expanded','false');
    btn.addEventListener('click', ()=>{
      const expanded = btn.getAttribute('aria-expanded') === 'true';
      btn.setAttribute('aria-expanded', String(!expanded));
      more.classList.toggle('hidden');
      btn.textContent = expanded ? 'More ▾' : 'Less ▴';
    });
    opts.appendChild(more); opts.appendChild(btn);
  }
  el.appendChild(opts);
  return el;
}

function optionNode(opt){
  const r = document.createElement('div'); r.className='opt';
  const f = document.createElement('span'); f.className='flag'; f.textContent = opt.split(' ')[0];
  const t = document.createElement('span'); t.textContent = opt;
  r.appendChild(f); r.appendChild(t);
  return r;
}

function filter(q){
  q = q.trim().toLowerCase();
  if(!q) return commands;
  return commands.filter(c => {
    if(c.name.includes(q)|| (c.alias && c.alias.includes(q))) return true;
    if(c.desc.toLowerCase().includes(q)) return true;
    return c.options.some(o=>o.toLowerCase().includes(q));
  });
}

searchInput.addEventListener('input', (e)=> renderCards(filter(e.target.value)));

// initial render: most commonly used at top (heuristic order preserved)
renderCards(commands);
